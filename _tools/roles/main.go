// The doc command prints the doc comment of a package-level object.
package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/loader"
)

const (
	// UASTPackage  package containing the roles definition.
	UASTPackage = "gopkg.in/bblfsh/sdk.v1/uast"
	// RoleType go type name of the Role type
	RoleType = UASTPackage + ".Role"
	// GitHubFilePattern route to the annotation.go file at GitHub
	GitHubFilePattern = "https://github.com/bblfsh/%s-driver/blob/master/driver/normalizer/annotation.go"
)

var (
	// OfficialDriver list of official driver maintanained by bblfsh org.
	OfficialDriver = map[string]string{
		"python": "github.com/bblfsh/python-driver/driver/normalizer",
		"java":   "github.com/bblfsh/java-driver/driver/normalizer",
	}
)

func main() {
	roles, err := findRoles()
	if err != nil {
		panic(err)
	}

	for l, pkg := range OfficialDriver {
		findUsage(l, pkg, roles)
	}

	fmt.Println(roles)

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
	conf.Import(pkg)
	prog, err := conf.Load()
	if err != nil {
		return err
	}

	info := prog.Package(pkg)
	result := &types.Info{
		Uses: make(map[*ast.Ident]types.Object),
	}

	tconf := types.Config{Importer: importer.Default()}
	_, err = tconf.Check("", prog.Fset, info.Files, result)
	if err != nil {
		return err
	}

	for id, obj := range result.Uses {
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
	for lang := range OfficialDriver {
		list = append(list, strings.Title(lang))
	}

	w.WriteString("Role|" + strings.Join(list, "|") + "\n")
	w.WriteString(strings.Repeat("-|-", len(list)) + "\n")
}

func writeTableBody(w *bytes.Buffer, r Roles) {
	for _, role := range r {
		fmt.Fprintf(w, "[%s](#%s)", role.Name, strings.ToLower(role.Name))
		for lang := range OfficialDriver {
			var used string
			if role.IsUsedBy(lang) {
				used = "✓"
			}

			fmt.Fprintf(w, "|%s", used)
		}

		fmt.Fprintf(w, "\n")
	}

	fmt.Fprintf(w, "\n\n")
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

		fmt.Fprintf(w, "## %s\n\n%s\n**Supported by**: %s\n\n",
			role.Name, role.Doc, strings.Join(l, ", "),
		)
	}
}
