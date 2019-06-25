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
	//命令行参数flag中可以赋值的数据
	//以逗号分隔的类型名称列表;必须设定
	typeNames = flag.String("type", "", "comma-separated list of type names; must be set")
	//输出文件名;默认srcdir / <type> _string.go
	output = flag.String("output", "", "output file name; default srcdir/<type>_toolMethod.go")
	//从生成的常量名称修剪`prefix`
	trimprefix = flag.String("trimprefix", "", "trim the `prefix` from the generated constant names")
	//使用行注释文本作为打印文本
	linecomment = flag.Bool("linecomment", false, "use line comment text as printed text when present")
	//以逗号分隔的要应用的构建标记列表
	buildTags = flag.String("tags", "", "comma-separated list of build tags to apply")
)

//自定义使用提示
//用法是替换flags包的Usage函数
// Usage is a replacement usage function for the flags package.
func Usage() {
	_, _ = fmt.Fprintf(os.Stderr, "Usage of structToMap:\n")
	_, _ = fmt.Fprintf(os.Stderr, "\tstructToMap [flags] -type T [directory]\n")
	_, _ = fmt.Fprintf(os.Stderr, "\tstructToMap [flags] -type T files... # Must be a single package\n")
	_, _ = fmt.Fprintf(os.Stderr, "For more information, see:\n")
	_, _ = fmt.Fprintf(os.Stderr, "\t百度orhttp://godoc.org/golang.org/x/tools/cmd/stringer\n")
	_, _ = fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults() //print flags
}

// Generator holds the state of the analysis. Primarily used to buffer
//Generator保持分析状态。主要用于缓冲
// the output for format.Source.
//生成器
type Generator struct {
	buf bytes.Buffer // Accumulated output.累计输出
	pkg *Package     // Package we are scanning. //正在扫描的包

	trimPrefix  string //修改前缀
	lineComment bool   //行注释
}

//print
func (g *Generator) Printf(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(&g.buf, format, args...)
}

type Package struct {
	name  string
	defs  map[*ast.Ident]types.Object //标识符
	files []*File                     //包含文件
}

// File holds a single parsed file and associated data.
//文件包含单个已解析的文件和关联数据。
type File struct {
	pkg     *Package  // Package to which this file belongs.此文件所属的包
	astFile *ast.File // Parsed AST.解析AST。
	// These fields are reset for each type being generated.
	//为每个生成的类型重置这些字段。
	typeName string // Name of the struct type.结构体类型的名称。
	//values   []Value // Accumulator for constant values of that type.累加器用于该类型的常量值。
	values      map[string]interface{}
	trimPrefix  string
	lineComment bool
}

// Structure field data
/*type Value struct {
	data map[string]interface{}
}
*/
//报告其指定的文件是否是目录
// isDirectory reports whether the named file is a directory.
func isDirectory(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
	}
	return info.IsDir() //true or false
}

// parsePackage analyzes the single package constructed from the patterns and tags.
//parsePackage分析从模式和标签构造的单个包。
// parsePackage exits if there is an error.
//如果出现错误，则退出parsePackage。
//解析当前package
func (g *Generator) parsePackage(patterns []string, tags []string) {
	cfg := &packages.Config{ //cfg:当前包的配置信息
		Mode: packages.LoadSyntax,
		// TODO: Need to think about constants in test files. Maybe write type_string_test.go
		// in a separate pass? For later.
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
//addPackage将类型检查的Package及其语法文件添加到生成器。
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
//generate为命名类型生成String方法。
func (g *Generator) generate(typeName string) {
	//values := make([]Value, 0, 100)
	var values map[string]interface{}
	for _, file := range g.pkg.files { //range all *ast.file
		// Set the state for this run of the walker.
		//设置此运行步行器的状态。
		file.typeName = typeName
		//file.values = nil
		if file.astFile != nil {
			//ast.Inspect(file.astFile, file.genDecl)
			ast.Inspect(file.astFile, func(node ast.Node) bool {
				// Find type
				ret, ok := node.(*ast.TypeSpec)
				if ok {
					//ret is structType?
					x, ok2 := (ret.Type).(*ast.StructType)
					if ok2 {
						//struct name
						if ret.Name.Name == file.typeName {
							//field list
							v := make(map[string]interface{})
							a := x.Fields.List
							for _, fielddata := range a { //a is sturuct fileds所有字段

								var fieldname string
								for _, j := range fielddata.Names { //name is i []  ??
									fieldname = j.Name
									data, ok := (fielddata.Type).(*ast.Ident)
									if ok {
										//ast.Ident
										v[fieldname] = data.Name
									} else {
										//ast.StarExpr
										data := (fielddata.Type).(*ast.StarExpr).X.(*ast.Ident)
										v[fieldname] = data.Name
									}

								}
								//匿名字段
								if fielddata.Names == nil {
									data := (fielddata.Type).(*ast.Ident)
									v[data.Name] = data.Name
								}
							}
							//println(v)
							//f.values = v
							//file.values=v
							values = v

						}
					}
					return true
				}
				return true
			})
			//values = file.values
		}
	}

	/*if len(values) == 0 {
		//运行stringer 没有参数 //typename (Pill)
		//找不到该类型
		log.Fatalf("no values defined for type %s", typeName)
	}*/
	if values == nil {
		//运行structToMap 没有参数 //typename (User)
		//找不到该类型
		log.Fatalf("no values defined for type %s", typeName)
	}

	g.buildOneRun(values, typeName)
}

// buildOneRun generates the variables and String method for a single run of contiguous values.
func (g *Generator) buildOneRun(data map[string]interface{}, typeName string) {

	g.Printf("\n")
	// The generated code is simple enough to write as a Printf format.
	//生成的代码非常简单，可以写为Printf格式。
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

// genDecl processes one declaration clause.
//genDecl处理一个声明子句。
func (f *File) genDecl(node ast.Node) bool {
	// Find type
	ret, ok := node.(*ast.TypeSpec)
	if ok {
		//ret is structType?
		x, ok2 := (ret.Type).(*ast.StructType)
		if ok2 {
			//struct name
			if ret.Name.Name == f.typeName {
				//field list
				v := make(map[string]interface{})
				a := x.Fields.List
				for _, fielddata := range a { //a is sturuct fileds所有字段

					var fieldname string
					for _, j := range fielddata.Names { //name is i []  ??
						fieldname = j.Name
						data, ok := (fielddata.Type).(*ast.Ident)
						if ok {
							//ast.Ident
							v[fieldname] = data.Name
						} else {
							//ast.StarExpr
							data := (fielddata.Type).(*ast.StarExpr).X.(*ast.Ident)
							v[fieldname] = data.Name
						}

					}
					//匿名字段
					if fielddata.Names == nil {
						data := (fielddata.Type).(*ast.Ident)
						v[data.Name] = data.Name
					}
				}
				//println(v)
				f.values = v

			}
		}
		return true
	}
	return true
}

// format returns the gofmt-ed contents of the Generator's buffer.
//format返回Generator缓冲区的gofmt-ed内容。
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
	flag.Usage = Usage             //使用提示，使用Usage替换原有的提示
	flag.Parse()                   //解析命令行参数
	if len(*typeNames) == 0 {
		//打印提示内容，退出程序
		flag.Usage()
		os.Exit(2)
	}
	//多个type参数时进行分割（window中用“”）
	types := strings.Split(*typeNames, ",")

	//构建标记
	var tags []string
	if len(*buildTags) > 0 {
		tags = strings.Split(*buildTags, ",")
	}

	// We accept either one directory or a list of files. Which do we have?
	//我们接受一个目录或一个文件列表。我们有哪些？
	//处理传入的文件名或者目录名
	//args:需要处理的目录或者文件
	args := flag.Args()
	if len(args) == 0 {
		// Default: process whole package in current directory.
		//默认值：处理当前目录中的整个包。
		args = []string{"."}
	}

	// Parse the package once.解析包一次。
	var dir string
	//创建一个生成器
	g := Generator{
		trimPrefix:  *trimprefix,  //前缀
		lineComment: *linecomment, //行注释
	}
	// TODO(suzmue): accept other patterns for packages (directories, list of files, import paths, etc).
	//接受包的其他模式（目录，文件列表，导入路径等）。
	//一次一个目录，且该目录必须存在
	if len(args) == 1 && isDirectory(args[0]) {
		dir = args[0]
	} else {
		if len(tags) != 0 {
			//-tags选项仅适用于目录，而不适用于指定文件的情况
			log.Fatal("-tags option applies only to directories, not when files are specified")
		}
		//Dir 返回路径的最后一个元素，通常是路径的目录。
		// 删除最后一个元素后，Dir 在路径上调用 Clean 并删除尾部斜线。
		// 如果路径为空，则 Dir 返回“。”。如果路径完全由分隔符组成，
		// 则 Dir 返回一个分隔符。返回的路径不会以分隔符结尾，除非它是根目录。
		dir = filepath.Dir(args[0])
	}

	g.parsePackage(args, tags)

	// Print the header and package clause.
	//打印header和package子句。
	g.Printf("// Code generated by \"structToMap %s\"; DO NOT EDIT.\n", strings.Join(os.Args[1:], " "))
	g.Printf("\n")
	g.Printf("package %s", g.pkg.name)
	g.Printf("\n")

	// Run generate for each type.
	//为每种类型运行生成method
	for _, typeName := range types {
		g.generate(typeName)
		//g.pkg.files=g.pkg.files[len(g.pkg.files):]
	}
	// Format the output.格式化输出
	src := g.format()

	// Write to file.
	outputName := *output
	if outputName == "" {
		baseName := fmt.Sprintf("%sToMap.go", types[0])
		outputName = filepath.Join(dir, strings.ToLower(baseName))
	}
	err := ioutil.WriteFile(outputName, src, 0644)
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}

}
