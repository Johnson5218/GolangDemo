package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// Define a template.
	const letter = `
您的快递预计 {{.Name}}前抵达
`

	// Prepare some data to insert into the template.
	type Recipient struct {
		Name string
	}
	var recipients = Recipient{
		"2019年12月16日17:18:18",
	}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(letter))

	// Execute the template for each recipient.
		err := t.Execute(os.Stdout, recipients)
		if err != nil {
			log.Println("executing template:", err)
		}

}
