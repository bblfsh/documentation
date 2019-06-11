package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sync"

	"github.com/bblfsh/sdk/v3/driver/manifest"
	"github.com/bblfsh/sdk/v3/driver/manifest/discovery"
	"github.com/heroku/docker-registry-client/registry"
)

const (
	org = discovery.GithubOrg
)

func main() {
	flag.Parse()
	if err := run(flag.Args()); err != nil {
		log.Fatal(err)
	}
}

func run(outFiles []string) error {
	ctx := context.TODO()
	langs, err := discovery.OfficialDrivers(ctx, &discovery.Options{
		NoStatic: true,
	})
	if err != nil {
		return err
	}
	names := make([]string, 0, len(langs))
	for _, d := range langs {
		names = append(names, d.Language)
	}
	log.Println(len(langs), "language drivers found:", names)

	ld := newLoader()

	var (
		list = make([]Driver, len(langs))

		wg sync.WaitGroup
		// limits the number of concurrent requests
		tokens = make(chan struct{}, 3)
		// first error
		errc = make(chan error, 1)
	)
	for i, d := range langs {
		list[i].Driver = d
		list[i].GithubURL = d.RepositoryURL()
		wg.Add(1)
		go func(d *Driver) {
			defer wg.Done()

			tokens <- struct{}{}
			defer func() {
				<-tokens
			}()

			if name := org + `/` + d.Language + `-driver`; ld.checkDockerImage(name) {
				d.DockerhubURL = `https://hub.docker.com/r/` + name + `/`
			}
			if vers, err := d.Versions(ctx); err == nil {
				d.Releases = vers
			} else {
				select {
				case errc <- fmt.Errorf("cannot fetch versions for %q: %v", d.Language, err):
				default:
				}
			}
		}(&list[i])
	}
	wg.Wait()
	select {
	case err := <-errc:
		return err
	default:
	}

	for _, fname := range outFiles {
		if err := writeFile(fname, list); err != nil {
			return err
		}
	}
	return nil
}

func writeFile(fname string, list []Driver) error {
	const filePerm = 0644
	switch filepath.Ext(fname) {
	case ".json":
		data, err := json.MarshalIndent(list, "", "\t")
		if err != nil {
			return err
		}
		// TODO(dennwc): A workaround for manifest encoding issue.
		//               We decided to use time.Time instead of *time.Time in SDK,
		//               and now ",omitempty" doesn't work properly.
		data = bytes.Replace(data, []byte(`
		"Build": "0001-01-01T00:00:00Z",`), nil, -1)
		return ioutil.WriteFile(fname, data, filePerm)
	case ".md":
		fallthrough
	default:
	}
	buf := bytes.NewBuffer(nil)

	buf.WriteString(header)

	buf.WriteString("\n## Supported languages\n")
	buf.WriteString(tableHeader)

	for _, m := range list {
		if !m.ForCurrentSDK() || m.InDevelopment() {
			continue
		}
		buf.WriteString(m.String())
	}

	written := false
	for _, m := range list {
		if m.ForCurrentSDK() && !m.InDevelopment() {
			continue
		}
		if !written {
			written = true
			buf.WriteString("\n## In development\n")
			buf.WriteString(tableHeader)
		}
		buf.WriteString(m.String())
	}

	buf.WriteString(footer)
	return ioutil.WriteFile(fname, buf.Bytes(), filePerm)
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

type Driver struct {
	discovery.Driver
	GithubURL    string              `json:",omitempty"`
	DockerhubURL string              `json:",omitempty"`
	Releases     []discovery.Version `json:",omitempty"`
}

func (m Driver) Maintainer() manifest.Maintainer {
	if len(m.Maintainers) == 0 {
		return manifest.Maintainer{Name: "-"}
	}
	return m.Maintainers[0]
}

const tableHeader = `
| Language   | Release | Status | SDK  | AST\* | Annotations\*\* | UAST\*\*\* | Container | Maintainer |
| :--------- | :------ | :----- | :--- | :--- | :------------- | :----- | :-------- | :--------- |
`

func (m Driver) String() string {
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
	latest := "-"
	if len(m.Releases) != 0 {
		vers := m.Releases[0].String()
		latest = link(vers, m.GithubURL+"/releases/tag/v"+vers)
	}
	return fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s | %s | %s |\n",
		link(name, m.GithubURL), latest, m.Status, m.SDKVersion,
		boolIcon(m.Supports(manifest.AST)),
		boolIcon(m.Supports(manifest.Roles)),
		boolIcon(m.Supports(manifest.UAST)),
		linkMark(m.DockerhubURL),
		link(mnt.Name, mlink),
	)
}

func (l *loader) checkDockerImage(name string) bool {
	// dockerhub site always returns 200, even if repository does not exists
	// so we will check image via Docker registry protocol
	m, err := l.r.Manifest(name, "latest")
	return err == nil && m != nil
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
# Languages
`

const footer = `
* \* The driver is able to return the native AST
* \*\* The driver is able to return the AST annotated
* \*\*\* The driver is able to return the UAST

**Don't see your favorite language? [Help us!](join-the-community.md)**
`
