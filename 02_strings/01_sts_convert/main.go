package main

import "fmt"

func main() {
	var a string
	var b string

	a = "hello world \n gogogo"
	b = `hello world
		\n
		gogogo`

	fmt.Println(a)
	fmt.Println(b)
}
