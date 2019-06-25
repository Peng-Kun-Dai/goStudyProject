package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/types"
	"golang.org/x/tools/go/packages"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	typeNames   = flag.String("type", "", "comma-separated list of type names; must be set")
	output      = flag.String("output", "", "output file name; default srcdir/<type>_toolMethod.go")
	trimprefix  = flag.String("trimprefix", "", "trim the `prefix` from the generated constant names")
	linecomment = flag.Bool("linecomment", false, "use line comment text as printed text when present")
	buildTags   = flag.String("tags", "", "comma-separated list of build tags to apply")
)

// Usage is a replacement usage function for the flags package.
func Usage() {
	_, _ = fmt.Fprintf(os.Stderr, "Usage of structToMap:\n")
	_, _ = fmt.Fprintf(os.Stderr, "\tstructToMapTool [flags] -type T [directory]\n")
	_, _ = fmt.Fprintf(os.Stderr, "\tstructToMapTool [flags] -type T files... # Must be a single package\n")
	_, _ = fmt.Fprintf(os.Stderr, "For more information, see:\n")
	_, _ = fmt.Fprintf(os.Stderr, "\thttps://www.jianshu.com/p/879edba89c72\n")
	_, _ = fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults() //print flags
}

// Generator holds the state of the analysis. Primarily used to buffer
// the output for format.Source.
type Generator struct {
	buf bytes.Buffer // Accumulated output.
	pkg *Package     // Package we are scanning.

	trimPrefix  string
	lineComment bool
}

type Package struct {
	name  string
	defs  map[*ast.Ident]types.Object
	files []*File
}

// File holds a single parsed file and associated data.
type File struct {
	pkg     *Package  // Package to which this file belongs.
	astFile *ast.File // Parsed AST.
	// These fields are reset for each type being generated.
	typeName    string // Name of the struct type.
	values      map[string]interface{}
	trimPrefix  string
	lineComment bool
}

//io writer
func (g *Generator) Printf(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(&g.buf, format, args...)
}

// isDirectory reports whether the named file is a directory.
func isDirectory(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
	}
	return info.IsDir() //true or false
}

//parsePackage analyzes the single package constructed from the patterns and tags.
// parsePackage exits if there is an error.
func (g *Generator) parsePackage(patterns []string, tags []string) {
	cfg := &packages.Config{
		Mode:       packages.LoadSyntax,
		Tests:      false,
		BuildFlags: []string{fmt.Sprintf("-tags=%s", strings.Join(tags, " "))},
	}
	pkgs, err := packages.Load(cfg, patterns...)
	if err != nil {
		log.Fatal(err)
	}
	if len(pkgs) != 1 {
		log.Fatalf("error: %d packages found", len(pkgs))
	}
	g.addPackage(pkgs[0])
}

// addPackage adds a type checked Package and its syntax files to the generator.
func (g *Generator) addPackage(pkg *packages.Package) {
	g.pkg = &Package{
		name:  pkg.Name,
		defs:  pkg.TypesInfo.Defs,
		files: make([]*File, len(pkg.Syntax)),
	}
	for i, file := range pkg.Syntax {
		g.pkg.files[i] = &File{
			astFile:     file,
			pkg:         g.pkg,
			trimPrefix:  g.trimPrefix,
			lineComment: g.lineComment,
		}
	}
}

// generate produces the String method for the named type.
func (g *Generator) generate(typeName string) {
	//save field information
	var values map[string]interface{}
	for _, file := range g.pkg.files { //range all *ast.file
		// Set the state for this run of the walker.
		file.typeName = typeName
		//file.values = nil
		if file.astFile != nil {
			ast.Inspect(file.astFile, func(node ast.Node) bool {
				// Find type
				ret, ok := node.(*ast.TypeSpec)
				if ok {
					//ret is structType? Correct name
					x, ok2 := (ret.Type).(*ast.StructType)
					if ok2 && (ret.Name.Name == file.typeName) {
						//field list
						v := make(map[string]interface{})
						a := x.Fields.List
						for _, fielddata := range a { //a is sturuct fileds
							var fieldname string
							for _, j := range fielddata.Names { //name is i []  ??
								fieldname = j.Name
								v[fieldname] = fielddata.Type
							}
							if fielddata.Names == nil {
								data := (fielddata.Type).(*ast.Ident)
								v[data.Name] = data.Name
							}
						}
						values = v
					}
					return true
				}
				return true
			})
		}
	}
	if values == nil {
		log.Fatalf("no values defined for type %s", typeName)
	}
	//code write into buf
	g.buildOneRun(values, typeName)
}

// buildOneRun generates the variables and String method for a single run of contiguous values.
func (g *Generator) buildOneRun(data map[string]interface{}, typeName string) {

	g.Printf("\n")
	// The generated code is simple enough to write as a Printf format.
	g.Printf(stringPrefix, typeName)
	for k := range data {
		g.Printf(stringCenter, k)
	}
	g.Printf(stringSuffix)
}

const stringPrefix = `func (i %[1]s) %[1]sToMap() map[string]interface{} {
	
	var data = make(map[string]interface{})
`
const stringCenter = `	data["%[1]s"] = i.%[1]s
`
const stringSuffix = `
	return data
}
`

// format returns the gofmt-ed contents of the Generator's buffer.
func (g *Generator) format() []byte {
	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		// Should never happen, but can arise when developing this code.
		//永远不应该发生，但在开发此代码时可能会出现。
		// The user can compile the output to see the error.
		//用户可以编译输出以查看错误。
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		return g.buf.Bytes()
	}
	return src
}
func main() {
	log.SetFlags(0)                //log output flag
	log.SetPrefix("structToMap: ") // log output Prefix
	flag.Usage = Usage
	flag.Parse()
	if len(*typeNames) == 0 {
		flag.Usage()
		os.Exit(2)
	}
	types := strings.Split(*typeNames, ",")

	var tags []string
	if len(*buildTags) > 0 {
		tags = strings.Split(*buildTags, ",")
	}

	// We accept either one directory or a list of files. Which do we have?
	args := flag.Args()
	if len(args) == 0 {
		// Default: process whole package in current directory.
		args = []string{"."}
	}

	// Parse the package once.
	var dir string
	g := Generator{
		trimPrefix:  *trimprefix,
		lineComment: *linecomment,
	}
	// TODO(suzmue): accept other patterns for packages (directories, list of files, import paths, etc).
	if len(args) == 1 && isDirectory(args[0]) {
		dir = args[0]
	} else {
		if len(tags) != 0 {
			log.Fatal("-tags option applies only to directories, not when files are specified")
		}
		dir = filepath.Dir(args[0])
	}
	g.parsePackage(args, tags)

	// Print the header and package clause.
	g.Printf("// Code generated by \"structToMap %s\"; DO NOT EDIT.\n\n", strings.Join(os.Args[1:], " "))
	g.Printf("package %s", g.pkg.name)
	g.Printf("\n")

	// Run generate for each type.
	for _, typeName := range types {
		g.generate(typeName)
	}
	// Format the output.
	src := g.format()

	// Write to file.
	outputName := *output
	if outputName == "" {
		baseName := fmt.Sprintf("%sToMap.go", strings.ToLower(types[0]))
		outputName = filepath.Join(dir, baseName)
	}
	err := ioutil.WriteFile(outputName, src, 0644)
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}

}
