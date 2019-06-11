// Generates a report about every driver wich UASTv2 types does it use,
// both in actual test fixtures and though the source code.
package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"sync"

	"github.com/bblfsh/sdk/v3/driver/manifest/discovery"
	"github.com/bblfsh/sdk/v3/uast"

	_ "net/http/pprof"
)

const (
	fixtureDir = "fixtures"
	fixtureExt = ".sem.uast"
	goDocURL   = "https://godoc.org/github.com/bblfsh/sdk/uast#"
	pprofAddr  = "localhost:6060"
)

var (
	reposRootPath = filepath.Join(".", "drivers")
	uastNodeRe    = regexp.MustCompile("uast:[^\"]*")

	pprof      = flag.Bool("pprof", false, "Start a pprof profiler HTTP service at "+pprofAddr)
	skipUpdate = flag.Bool("skip", false, "skip git clone or pull")
)

func main() {
	flag.Parse()

	if *pprof {
		fmt.Fprintf(os.Stderr, "running pprof on %s\n", pprofAddr)
		go func() {
			if err := http.ListenAndServe(pprofAddr, nil); err != nil {
				log.Fatalf("cannot start pprof: %v", err)
			}
		}()
	}

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	drivers, err := listDrivers()
	if err != nil {
		return fmt.Errorf("failed to list drivers: %s", err)
	}

	if !*skipUpdate {
		if err := maybeCloneOrPullAll(drivers); err != nil {
			return fmt.Errorf("failed to pull driver repos: %s", err)
		}
	}

	uastTypes := findUASTTypesInSDK()
	for _, driver := range drivers {
		if err := analyzeFixtures(driver); err != nil {
			return err
		}
		analyzeCode(driver, uastTypes)
	}

	formatMarkdownTable(drivers, uastTypes)
	return nil
}

type driverStats struct {
	url                 string
	language            string
	path                string         // path in local FS to the driver source code clone
	skip                bool           // skip including the driver to the final report
	fixtureCount        int            // number of fixture files for a driver
	uastInFixturesCount map[string]int // number of times UAST type used in fixtures
	uastInCodeCount     map[string]int // number of times UAST type used in code
}

// listDrivers lists all available drivers.
func listDrivers() ([]*driverStats, error) {
	fmt.Fprintf(os.Stderr, "discovering all available drivers\n")
	langs, err := discovery.OfficialDrivers(context.TODO(), &discovery.Options{
		NoStatic:      true,
		NoMaintainers: true,
		NoBuildInfo:   true,
	})
	if err != nil {
		return nil, err
	}

	drivers := make([]*driverStats, 0, len(langs))
	for _, l := range langs {
		if !l.ForCurrentSDK() || l.InDevelopment() {
			continue
		}
		drivers = append(drivers, &driverStats{
			language:            l.Language,
			url:                 l.RepositoryURL(),
			path:                l.RepositoryURL()[strings.LastIndex(l.RepositoryURL(), "/"):],
			uastInFixturesCount: make(map[string]int),
			uastInCodeCount:     make(map[string]int),
		})
	}
	fmt.Fprintf(os.Stderr, "%d drivers available on-line\n", len(langs))
	return drivers, nil
}

// maybeCloneOrPullAll either clones repos to path in local FS or, if already preset,
// pulls the latest master for each of them.
func maybeCloneOrPullAll(drivers []*driverStats) error {
	fmt.Fprintf(os.Stderr, "cloning %d drivers to %s\n", len(drivers), reposRootPath)
	err := os.MkdirAll(reposRootPath, os.ModePerm)
	if err != nil {
		return err
	}

	var (
		throttle = make(chan int, 3) // fetches to run in parallel
		wg       sync.WaitGroup
	)
	for i := range drivers {
		driver := drivers[i]
		wg.Add(1)
		go func() {
			throttle <- 1
			defer func() { <-throttle; wg.Done() }()
			maybeCloneOrPull(driver)
		}()
	}
	wg.Wait()
	return nil
}

func maybeCloneOrPull(d *driverStats) {
	repoPath := filepath.Join(reposRootPath, d.path)
	_, err := os.Stat(repoPath)
	if os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "%s does not exist, cloning from %s\n", repoPath, d.url)
		cmd := exec.Command("git", "clone", d.url+".git")
		cmd.Dir = reposRootPath
		err = cmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to git clone %s\n", d.url)
		}
	} else if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read %q: %s\n", repoPath, err)
	}

	fmt.Fprintf(os.Stderr, "%s dir exists, will git pull instead\n", repoPath)
	cmd := exec.Command("git", "pull", "origin", "master")
	cmd.Dir = repoPath
	err = cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed git pull %s\n", d.url)
	}
}

type uastType struct {
	name string
	// TODO decide what else is needed for analyzeCode()
	// or replace struct with string
}

// isUsedIn checks if current UAST type is used in the given package.
func (u *uastType) isUsedIn() bool {
	// TODO implement for analyzeCode()
	return false
}

// findUASTTypesInSDK finds all types from SDK.
func findUASTTypesInSDK() []uastType {
	var out []uastType
	// TODO: either parse code of the package, iterate all structs
	// or expose all registered types though SDK and use it here
	types := []interface{}{
		uast.Position{},
		uast.Positions{},
		uast.Identifier{},
		uast.String{},
		uast.Bool{},
		uast.QualifiedIdentifier{},
		uast.Comment{},
		uast.Group{},
		uast.FunctionGroup{},
		uast.Block{},
		uast.Alias{},
		uast.Import{},
		uast.RuntimeImport{},
		uast.RuntimeReImport{},
		uast.InlineImport{},
		uast.Argument{},
		uast.FunctionType{},
		uast.Function{},
	}
	for _, typee := range types {
		out = append(out, uastType{reflect.TypeOf(typee).String()})
	}
	fmt.Fprintf(os.Stderr, "%d uast:* types found in SDK\n", len(out))
	return out
}

// analyzeFixtures goes though all fixtures, assuming the driver is cloned.
// It updates given driverStats with results.
func analyzeFixtures(driver *driverStats) error {
	fixDir := filepath.Join(reposRootPath, driver.path, fixtureDir)
	fixtureFiles, err := lsDir(fixDir)
	if os.IsNotExist(err) {
		driver.skip = true
		return nil
	} else if err != nil {
		return err
	}
	driver.fixtureCount += len(fixtureFiles)

	for _, file := range fixtureFiles {
		if !strings.HasSuffix(file.Name(), fixtureExt) {
			continue
		}

		data, err := ioutil.ReadFile(filepath.Join(fixDir, file.Name()))
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to read %q, skipping\n", file.Name())
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			if !strings.Contains(line, "uast:") {
				continue
			}
			for _, uastType := range uastNodeRe.FindAllString(line, -1) {
				driver.uastInFixturesCount[strings.Replace(uastType, ":", ".", 1)]++
			}
		}
	}
	return nil
}

// lsDir lists all files in the given dir.
// It does not use ioutil.ReadDir as we do not care about files order.
func lsDir(dir string) ([]os.FileInfo, error) {
	fmt.Fprintf(os.Stderr, "reading %s/*%s files\n", dir, fixtureExt)
	f, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	return files, nil
}

// analyzeCode checks if any of the types are used by
// this driver's package, though analyzing it's AST.
func analyzeCode(driver *driverStats, uasts []uastType) {
	// TODO:
	// load package
	// for _, typee := range uasts {
	//   if typee.isUsedIn(package) {
	//     driver.codeUast[typee]++
	//   }
	// }
}

func formatMarkdownTable(drivers []*driverStats, uastTypes []uastType) {
	fmt.Print(header)
	defer fmt.Print(footer)

	allDrivers := len(drivers)
	drs := drivers[:0] // filter drivers \wo fixtures
	for _, x := range drivers {
		if !x.skip {
			drs = append(drs, x)
		}
	}
	fmt.Fprintf(os.Stderr, "only %d drivers has fixtures, out of %d\n", len(drs), allDrivers)

	formatMarkdownTableHeader(drs)
	for _, typee := range uastTypes {
		// %25s produces nice ASCII result
		fmt.Printf("|[%s](%s)|",
			strings.Replace(typee.name, ".", ":", 1),
			goDocURL+typee.name[strings.IndexRune(typee.name, '.')+1:],
		)
		for _, dr := range drs {
			if dr.uastInCodeCount[typee.name] > 0 {
				fmt.Printf(" %d/%d |",
					dr.uastInFixturesCount[typee.name],
					dr.uastInCodeCount[typee.name],
				)
			} else {
				fmt.Printf(" %d |", dr.uastInFixturesCount[typee.name])
			}
		}
		fmt.Println()
	}
}

func formatMarkdownTableHeader(drivers []*driverStats) {
	fmt.Printf("|%25s|", "")
	for _, dr := range drivers {
		// %5s produces nice ASCII result
		fmt.Printf("[%s](%s)|", dr.language, dr.url)
	}
	fmt.Print("\n| :---------------------- |")
	for range drivers {
		fmt.Printf(" :-- |")
	}
	fmt.Println()
}

const header = `<!-- Code generated by 'make types' DO NOT EDIT. -->
# UAST Types

For every [UAST type](uast/semantic-uast.md#types)
in every driver the following two values are reported:
 - _fixtures usage_  - number of times this type was used in driver _fixtures_ (_*.sem.uast_ files)
 - _code usage_ - number of times this type was usind in the driver mapping DSL code (_normalizer.go_ file)

The format is <_fixtures usage_>/<_code usage_> in case _code usage_ is not zero.
Otherwise, only _fixture usage_ is report.

`

const footer = `
**Don't see your favorite AST construct represented? [Help us!](join-the-community.md)**
`
