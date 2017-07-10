package main

import (
	"html/template"
	"log"
	"os"
)

type Note struct {
	Title       string
	Description string
}

//const tmpl = `Note - Title {{.Title}}, Description - {{.Description}}`
const tmplRange = `Notes are :
	{{range.}}
		Title : {{.Title}}, Description - {{.Description}}
	{{end}}
`

func main() {
	note := []Note{
		{"text/templates", "Templates generate textual output"},
		{"html/templates", "Templates generate html output"},
	}
	t := template.New("note")
	t, err := t.Parse(tmplRange)
	if err != nil {
		log.Fatal("Parse :", err)
		return
	}

	if err := t.Execute(os.Stdout, note); err != nil {
		log.Fatal("Execute : ", err)
		return
	}
}
