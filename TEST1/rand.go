package main

import "fmt"
import "math/rand"


func main() {
	r := rand.New(rand.NewSource(99))
	var temp = r.Int31()
	for i := 0; i < 20; i++ {
		temp = r.Int31()
		fmt.Println(temp)
	}
	
		fmt.Println("**********************")
		fmt.Println(temp)
}
