package main

import "fmt"
import "strconv"
import "math/rand"


func main() {
	var temp = "template_"
	temp+=strconv.Itoa(001)
	fmt.Println(temp)
	r := rand.New(rand.NewSource(99))

	for i := 0; i < 20; i++ {
		fmt.Println(r.Int31())
	}
}
