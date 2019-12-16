package main

import "fmt"


func sum(str string, options ...int) (x int) {
	for _, y := range options {
		x += y
	}
	fmt.Println(str)
	return 
}

func main() {
	a := sum("add all", 1, 2, 3, 4, 5)
	fmt.Println(a)
}
