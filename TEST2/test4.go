package main

import (
	"fmt"
	"strconv"
	"bytes"
	"text/template"
)

func main() {
	var str = [3]string{"go核心编程", "12", "2019-12-18 17:28:10"}
	var m1 map[string]string 
	m1 = make(map[string]string)
	for i, index := range str {
		temp := "STR" temp += strconv.Itoa(i+1) m1[temp] = index
	}
	letter := `
	您的快递 {{.STR1}} 距离您只有 {{.STR2}} 公里了，预计 {{.STR3}} 到达
	`
	t := template.Must(template.New("letter").Parse(letter))
	// Execute the template for each recipient.
	buf := new(bytes.Buffer)
	err := t.Execute(buf, m1)
	fmt.Println(buf)
		// import "bytes"
	if err != nil {
		buf.String() // returns a string of what was written to it
		fmt.Println(buf)
	} 
}

