package main

func a(i int) func() {
	return func() {
		println(i)
	}
}


func b(i int) func () {
	return func(){
		print(i+1)
	}
}


func main() {
	f := a(1)
	f()

	d := b(2)
	d()
}
