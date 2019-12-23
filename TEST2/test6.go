package main

func a(i int) func() {
	return func() {
		println(i)
	}
}

func main() {
	f := a(1)
	f()
}
