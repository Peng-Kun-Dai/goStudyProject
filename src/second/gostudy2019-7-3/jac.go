package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"text/template"
)

const signleFuncTemplate = `
func (i {{.StructName}}) {{.StructName}}ToMap() map[string]interface{} {
	var data = make(map[string]interface{})
	{{range .FieldsNames}}
	data["{{.}}"] = i.{{.}}
	{{end}}
	return data
}
`
const letter = `
//commandline
`

type gen struct {
	StructName  string
	FieldsNames []string
}

func main() {
	g := gen{
		StructName:  "User",
		FieldsNames: []string{"name", "age"},
	}
	//t := template.Must(template.New("letter").Parse(letter))
	//t := template.Must(template.New("functemplate").Parse(signleFuncTemplate))
	t, err := template.ParseFiles("./filetmp/file1.tmpl")
	if err != nil {
		log.Println("parseFilexxx:", err)
	}
	var f = bytes.Buffer{}
	err = t.Execute(&f, g)
	if err != nil {
		log.Println("executing template:", err)
	}
	err = ioutil.WriteFile("x.go", f.Bytes(), 0644)
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}

}
