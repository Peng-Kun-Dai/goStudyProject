package test

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/scanner"
	"go/token"
	"log"
	"os"
	"testing"
	"text/template"
)

func TestScanner(t *testing.T) {
	src := []byte(`package main
type User struct {
	name string
	age  int
}
type date struct {
	name string
	age  string
}
`)
	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, src, nil, 0)

	for {
		pos, tok, lit := s.Scan()
		fmt.Printf("%-6s%-8s%q\n", fset.Position(pos), tok, lit)

		if tok == token.EOF {
			break
		}
	}
}
func TestInspectAST(t *testing.T) {
	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "test2.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	//Store all structure information
	structdata := make(map[string]interface{})
	ast.Inspect(f, func(n ast.Node) bool {
		// Find type
		ret, ok := n.(*ast.TypeSpec)
		if ok {
			//ret is structType?
			x, ok2 := (ret.Type).(*ast.StructType)
			if ok2 {
				//struct name
				structName := ret.Name.Name
				//field list
				a := x.Fields.List
				//Store a structure information
				everyStructFieldData := make(map[string]interface{})
				for _, fielddata := range a { //a is sturuct fileds所有字段

					var fieldname string
					for _, j := range fielddata.Names { //name is i []  ??
						fieldname = j.Name
						data, ok := (fielddata.Type).(*ast.Ident)
						if ok {
							//ast.Ident
							everyStructFieldData[fieldname] = data.Name
						} else {
							//ast.StarExpr
							data := (fielddata.Type).(*ast.StarExpr).X.(*ast.Ident)
							everyStructFieldData[fieldname] = data.Name
						}

					}
					//匿名字段
					if fielddata.Names == nil {
						data := (fielddata.Type).(*ast.Ident)
						everyStructFieldData[data.Name] = data.Name
					}
				}
				structdata[structName] = everyStructFieldData
			}
			return true
		}
		return true

		/*// Find Return Statements
		ret2,ok2:=n.(*ast.StructType)  //all structType
		if ok2 {
			fmt.Println(ret2.Struct)
			fmt.Println()
			//fmt.Println(ret.Type)
			//fmt.Println(ret.Fields.NumFields())  //fileds
			fmt.Printf("struct found on line %v:\n", fset.Position(ret2.Pos())) //pos
			printer.Fprint(os.Stdout, fset, ret2)
			fmt.Printf("\n")
			return true
		}
		return true*/
	})
	fmt.Println(structdata)
}

const functemplate = `

func (i {{.structName}}) {{.structName}}ToMap() map[string]interface{} {
    var data = make(map[string]interface{})
    {{range .fieldsNames.Field}}
    data["{{.Name}}"] = i.{{.Name}}
    return data
}
`

type gen struct {
	structName  string
	fieldsNames []string
}

func testTem() {
	g := gen{
		structName:  "jayce",
		fieldsNames: []string{"User", "age"},
	}
	t := template.Must(template.New("letter").Parse(functemplate))
	err := t.Execute(os.Stdout, g)
	if err != nil {
		log.Println("executing template:", err)
	}

}
