package main

import(
	"fmt"
	"strconv"
)

func main() {
	var str string
	str = ""
	temp, err := strconv.ParseInt(str, 10, 64)
	fmt.Println(temp, err)
}
