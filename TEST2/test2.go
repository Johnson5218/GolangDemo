package main

import (
	"bytes"
	"text/template"
	"fmt"
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
	fmt.Println(letter)	
	buf := new(bytes.Buffer)
	err := t.Execute(buf, recipients)
		// import "bytes"
	buf.String() // returns a string of what was written to it
	fmt.Println(buf)
	fmt.Println(err)
}
