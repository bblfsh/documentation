// The doc command prints the doc comment of a package-level object.
package main

import (
	"bytes"
	"context"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"sort"
	"strings"

	"golang.org/x/tools/go/loader"
	"gopkg.in/bblfsh/sdk.v2/manifest/discovery"
)

const (
	// UASTPackage  package containing the roles definition.
	UASTPackage = "gopkg.in/bblfsh/sdk.v2/uast"
	// RoleType go type name of the Role type
	RoleType = UASTPackage + ".Role"
	// GitHubFilePattern route to the annotation.go file at GitHub
	GitHubFilePattern = "https://github.com/bblfsh/%s-driver/blob/master/driver/normalizer/annotation.go"
)

var (
	// OfficialDriver list of official driver maintanained by bblfsh org.
	OfficialDriver []discovery.Driver
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	roles, err := findRoles()
	if err != nil {
		return err
	}

	list, err := discovery.OfficialDrivers(context.TODO(), &discovery.Options{
		NoMaintainers: true,
	})
	if err != nil {
		return err
	}
	for i := 0; i < len(list); i++ {
		if !list[i].IsRecommended() {
			list = append(list[:i], list[i+1:]...)
			i--
		}
	}
	OfficialDriver = list

	var last error
	for _, d := range OfficialDriver {
		if !d.IsRecommended() {
			continue
		}
		pkg := strings.TrimPrefix(d.RepositoryURL(), "https://")
		pkg += "/driver/normalizer"
		if err := findUsage(d.Language, pkg, roles); err != nil {
			log.Println(err)
			last = err
		}
	}

	fmt.Println(roles)
	return last
}

// findRoles find the roles defined at the uast package.
func findRoles() (Roles, error) {
	var out Roles

	conf := loader.Config{ParserMode: parser.ParseComments}
	conf.Import(UASTPackage)
	prog, err := conf.Load()
	if err != nil {
		return nil, err
	}

	pkg := prog.Package(UASTPackage).Pkg
	for _, name := range pkg.Scope().Names() {
		obj := pkg.Scope().Lookup(name)
		if obj.Type().String() != RoleType {
			continue
		}

		out = append(out, &Role{
			Name:      obj.Name(),
			Doc:       findDoc(prog, obj.Pos()).Text(),
			Languages: make(map[string][]token.Position),
		})
	}

	return out, nil
}

func findDoc(prog *loader.Program, pos token.Pos) *ast.CommentGroup {
	_, path, _ := prog.PathEnclosingInterval(pos, pos)
	for _, n := range path {
		n, ok := n.(*ast.ValueSpec)
		if !ok {
			continue
		}

		return n.Doc
	}

	return nil
}

// findUsage finds in the normalizer package of a driver which roles are being
// used.
func findUsage(language, pkg string, roles Roles) error {
	conf := loader.Config{ParserMode: parser.ParseComments}
	conf.Import(pkg) // TODO: clone it if not found
	prog, err := conf.Load()
	if err != nil {
		return err
	}

	info := prog.Package(pkg)
	for id, obj := range info.Uses {
		if obj.Type().String() != RoleType {
			continue
		}

		roles.UsedBy(obj.Name(), language, prog.Fset.Position(id.Pos()))
	}

	return nil

}

// Role contains the relevant information of a Role definition
type Role struct {
	Name      string
	Doc       string
	Languages map[string][]token.Position
}

func (r *Role) IsUsedBy(language string) bool {
	l, ok := r.Languages[language]
	if ok && len(l) > 0 {
		return true
	}

	return false
}

// Roles is a list of roles.
type Roles []*Role

// UsedBy adds the given language to the list of language using a specific role.
func (r Roles) UsedBy(name, language string, pos token.Position) {
	for _, role := range r {
		if role.Name != name {
			continue
		}

		if _, ok := role.Languages[language]; !ok {
			role.Languages[language] = make([]token.Position, 0)
		}

		role.Languages[language] = append(role.Languages[language], pos)
	}
}

const documentHeader = "" +
	"# Roles list\n\n" +
	"Role is the main UAST annotation. It indicates that a node in an AST " +
	"can be interpreted as acting with certain language-independent role.\n\n"

func (r Roles) String() string {
	buf := bytes.NewBuffer([]byte(documentHeader))
	writeTableHeader(buf)
	writeTableBody(buf, r)
	writeList(buf, r)

	return buf.String()
}

func writeTableHeader(w *bytes.Buffer) {
	var list []string
	for _, d := range OfficialDriver {
		list = append(list, d.Name)
	}

	w.WriteString("Role |" + strings.Join(list, " | ") + "\n")
	w.WriteString("---" + strings.Repeat("|---", len(list)) + "\n")
}

func writeTableBody(w *bytes.Buffer, r Roles) {
	for _, role := range r {
		fmt.Fprintf(w, "[%s](#%s) ", role.Name, strings.ToLower(role.Name))
		for _, d := range OfficialDriver {
			var used string
			if role.IsUsedBy(d.Language) {
				used = "✓"
			}

			fmt.Fprintf(w, " | %s", used)
		}

		fmt.Fprint(w, "\n")
	}

	fmt.Fprint(w, "\n\n")
}

func writeList(w *bytes.Buffer, r Roles) {
	for _, role := range r {

		var l []string
		for language := range role.Languages {
			l = append(l, fmt.Sprintf("[*%s*](%s)",
				strings.Title(language),
				fmt.Sprintf(GitHubFilePattern, language),
			))
		}
		sort.Strings(l)

		fmt.Fprintf(w, "## %s\n\n%s\n**Supported by**: %s\n\n",
			role.Name, role.Doc, strings.Join(l, ", "),
		)
	}
}
