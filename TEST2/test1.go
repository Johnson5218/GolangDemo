package main

import (
	"log"
	"os"
	"text/template"
)
func main() {
	// Define a template.
	const letter = `
{{with .Str1 -}}
你的快递预计 {{.}}到达.
{{end}}
`

	// Prepare some data to insert into the template.
	type Recipient struct {
		Str1 string
	}
	var recipients = Recipient{"2019年12月16日17:00:55"}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(letter))

	// Execute the template for each recipient.
		err := t.Execute(os.Stdout, recipients)
		if err != nil {
			log.Println("executing template:", err)
		}

}
