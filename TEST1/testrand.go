package main

import "fmt"
import "math/rand"
import "time"

func main() {
	//r := rand.New(rand.Seed(time.Now().Unix()))
	//num := r.Int31()
	num := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000000)
	fmt.Println(num)
}
