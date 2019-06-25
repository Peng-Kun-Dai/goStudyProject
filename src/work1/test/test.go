package test

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	structImformation := findOnUser()
	fmt.Println(structImformation)
}
func getStrcutFromFile() map[string]interface{} {
	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "struct.go", nil, parser.ParseComments)
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
	return structdata
}

// only find User
func findOnUser() map[string]interface{} {
	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "struct.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	//Store a structure information
	everyStructFieldData := make(map[string]interface{})
	ast.Inspect(f, func(n ast.Node) bool {
		// Find type
		ret, ok := n.(*ast.TypeSpec)
		if ok {
			//ret is structType?
			x, ok2 := (ret.Type).(*ast.StructType)
			if ok2 {
				//struct name
				if ret.Name.Name == "User" {
					//field list
					a := x.Fields.List

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

				}
			}
			return true
		}
		return true
	})
	return everyStructFieldData
}
