package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/heroku/docker-registry-client/registry"
	"gopkg.in/bblfsh/sdk.v1/manifest"
)

const (
	org            = "bblfsh"
	driverBlobsURL = "https://raw.githubusercontent.com/" + org + "/%s-driver/master/"
	manifestURL    = driverBlobsURL + "manifest.toml"
	maintainersURL = driverBlobsURL + "MAINTAINERS"
)

func main() {
	if err := run(os.Stdout); err != nil {
		log.Fatal(err)
	}
}

func run(w io.Writer) error {
	langs, err := getDriverLanguages()
	if err != nil {
		return err
	}
	log.Println(len(langs), "language drivers found:", langs)

	ld := newLoader()
	fmt.Fprint(w, header)
	defer fmt.Fprint(w, footer)

	var (
		wg   sync.WaitGroup
		mu   sync.Mutex
		list []*Manifest
		last error
		// limits the number of concurrent requests
		tokens = make(chan struct{}, 3)
	)
	for _, id := range langs {
		id := id
		wg.Add(1)
		go func() {
			defer wg.Done()

			tokens <- struct{}{}
			defer func() {
				<-tokens
			}()

			m, err := ld.getManifest(id)
			if err != nil {
				mu.Lock()
				last = err
				mu.Unlock()
				log.Println(id, err)
			} else {
				mu.Lock()
				list = append(list, m)
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	// sort by status, festures, name
	sort.Slice(list, func(i, j int) bool {
		a, b := list[i], list[j]
		if s1, s2 := statuses[a.Status], statuses[b.Status]; s1 > s2 {
			return true
		} else if s1 < s2 {
			return false
		}
		if n1, n2 := len(a.Features), len(b.Features); n1 > n2 {
			return true
		} else if n1 < n2 {
			return false
		}
		return a.Language < b.Language
	})

	fmt.Fprintln(w, "\n# Supported languages")
	fmt.Fprint(w, tableHeader)

	li := len(list)
	for i, m := range list {
		if statuses[m.Status] < statuses[manifest.Alpha] {
			li = i
			break
		}
		fmt.Fprint(w, m.String())
	}

	list = list[li:]
	if len(list) == 0 {
		return last
	}

	fmt.Fprintln(w, "\n# In development")
	fmt.Fprint(w, tableHeader)

	for _, m := range list {
		fmt.Fprint(w, m.String())
	}

	return last
}

var statuses = map[manifest.DevelopmentStatus]int{
	manifest.Inactive: 0,
	manifest.Planning: 1,
	manifest.PreAlpha: 2,
	manifest.Alpha:    3,
	manifest.Beta:     4,
	manifest.Stable:   5,
	manifest.Mature:   6,
}

var reDriverLink = regexp.MustCompile(`href="/?` + org + `/([^\s-]+)-driver`)

func getDriversPage(n int) ([]string, error) {
	const base = `https://github.com/search`

	par := make(url.Values)
	par.Set("type", "Repositories")
	par.Set("utf8", "✓")
	par.Set("q", "topic:driver topic:babelfish org:"+org)
	par.Set("p", strconv.Itoa(n))

	resp, err := http.Get(base + "?" + par.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("status: %v", resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	sub := reDriverLink.FindAllSubmatch(data, -1)
	seen := make(map[string]struct{})
	for _, s := range sub {
		seen[string(s[1])] = struct{}{}
	}
	var out []string
	for s := range seen {
		out = append(out, s)
	}
	return out, nil
}

func getDriverLanguages() ([]string, error) {
	var out []string
	for i := 1; ; i++ {
		page, err := getDriversPage(i)
		if err != nil {
			return out, err
		}
		out = append(out, page...)
		if len(page) == 0 {
			break
		}
	}
	return out, nil
}

func newLoader() *loader {
	r, err := registry.New("https://registry-1.docker.io/", "", "")
	if err != nil {
		panic(err)
	}
	return &loader{r: r}
}

type loader struct {
	r *registry.Registry
}

func (l *loader) getManifest(id string) (*Manifest, error) {
	resp, err := http.Get(fmt.Sprintf(manifestURL, id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	name := org + `/` + id + `-driver`
	gh := `https://github.com/` + name
	if resp.StatusCode == http.StatusNotFound {
		return &Manifest{
			Manifest: manifest.Manifest{
				Language: id,
				Status:   manifest.Inactive,
			},
			GithubURL: gh,
		}, nil
	} else if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("status: %v", resp.Status)
	}
	var m manifest.Manifest
	if err := m.Decode(resp.Body); err != nil {
		return nil, err
	}
	mf := &Manifest{
		Manifest:     m,
		GithubURL:    gh,
		DockerhubURL: `https://hub.docker.com/r/` + name + `/`,
	}
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		if !checkURL(mf.GithubURL) {
			mf.GithubURL = ""
		}
	}()
	go func() {
		defer wg.Done()
		if !l.checkDockerImage(name) {
			mf.DockerhubURL = ""
		}
	}()
	go func() {
		defer wg.Done()
		mf.Maintainers = getMaintainers(mf.Language)
	}()
	wg.Wait()
	return mf, nil
}

type Maintainer struct {
	Name   string
	Email  string
	Github string
}

type Manifest struct {
	manifest.Manifest
	GithubURL    string
	DockerhubURL string
	Maintainers  []Maintainer
}

func (m Manifest) Maintainer() Maintainer {
	if len(m.Maintainers) == 0 {
		return Maintainer{Name: "-"}
	}
	return m.Maintainers[0]
}

func (m Manifest) String() string {
	name := m.Name
	if name == "" {
		name = m.Language
	}
	mnt := m.Maintainer()
	var mlink string
	if mnt.Github != "" {
		mnt.Name = mnt.Github
		mlink = `https://github.com/` + mnt.Github
	} else if mnt.Email != "" {
		mlink = `mailto:` + mnt.Email
	}
	return fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s | %s |\n",
		link(name, m.GithubURL), m.Language, m.Status,
		boolIcon(m.HasFeature(manifest.AST)),
		boolIcon(m.HasFeature(manifest.UAST)),
		boolIcon(m.HasFeature(manifest.Roles)),
		linkMark(m.DockerhubURL),
		link(mnt.Name, mlink),
	)
}
func (m Manifest) HasFeature(f manifest.Feature) bool {
	for _, f2 := range m.Features {
		if f == f2 {
			return true
		}
	}
	return false
}
func checkURL(url string) bool {
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode/100 == 2
}
func (l *loader) checkDockerImage(name string) bool {
	// dockerhub site always returns 200, even if repository does not exists
	// so we will check image via Docker registry protocol
	m, err := l.r.Manifest(name, "latest")
	return err == nil && m != nil
}

var reMaintainer = regexp.MustCompile(`^([^<(]+)\s<([^>]+)>(\s\(@([^\s]+)\))?`)

func getMaintainers(id string) []Maintainer {
	resp, err := http.Get(fmt.Sprintf(maintainersURL, id))
	if err != nil {
		log.Println(err)
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode/100 != 2 {
		return nil
	}
	var out []Maintainer
	sc := bufio.NewScanner(resp.Body)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		sub := reMaintainer.FindStringSubmatch(line)
		if len(sub) == 0 {
			continue
		}
		m := Maintainer{Name: sub[1], Email: sub[2]}
		if len(sub) >= 5 {
			m.Github = sub[4]
		}
		out = append(out, m)
	}
	return out
}

func boolIcon(v bool) string {
	if v {
		return "✓"
	}
	return "✗"
}

func linkMark(url string) string {
	if url == "" {
		return boolIcon(false)
	}
	return link(boolIcon(true), url)
}

func link(name, url string) string {
	if url == "" {
		return name
	}
	return fmt.Sprintf(`[%s](%s)`, name, url)
}

const header = `<!-- Code generated by 'make languages' DO NOT EDIT. -->
`

const tableHeader = `
| Language   | Key        | Status  | AST\* | UAST\*\* | Annotations\*\*\* | Container | Maintainer |
| ---------- | ---------- | ------- | ---- | ------ | -------------- | --------- | ---------- |
`

const footer = `
- \* The driver is able to return the native AST
- \*\* The driver is able to return the UAST
- \*\*\* The driver is able to return the UAST annotated


**Don't see your favorite language? [Help us!](community.md)**
`
