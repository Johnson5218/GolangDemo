package main
import "fmt"
import "math/rand"
import "time"
func main() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := rnd.Int31n(1000000)
	fmt.Println(vcode)
}
