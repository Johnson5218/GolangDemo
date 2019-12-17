package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	// Define a template.
	/*
	content := `
Dear {{.Name}},
{{if .Attended}}
It was a pleasure to see you at the wedding.
{{- else}}
It is a t you couldn't make it to the wedding.
{{- end}}
{{with .Gift -}}
Thank you for the lovely {{.}}.
{{end}}
Best wishes,
Josie
`

	// Prepare some data to insert into the template.
	type Recipient struct {
		Name, Gift string
		Attended   bool
	}
*/
letter := `您的快递预计{{ .Str1 }}到达`
	
	type Rec struct {
		Str1 string
	}
	rec := Rec{"2019年12月16日16:53:07"}

	t := template.Must(template.New("letter").Parse(letter))
	err := t.Execute(os.Stdout, rec)
	if err != nil {
		fmt.Println(err)
	}
	
/*
	var recipients = []Recipient{
		{"Aunt Mildred", "bone china tea set", true},
		{"Uncle John", "moleskin pants", false},
		{"Cousin Rodney", "", false},
	}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("content").Parse(content))

	// Execute the template for each recipient.
	for _, r := range recipients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}
*/
}
