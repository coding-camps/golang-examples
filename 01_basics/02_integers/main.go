package main

import "fmt"

func main() {

	var a int = 10
	fmt.Printf("a = %d\n", a)
	a++
	fmt.Printf("b = %d\n", a)
	a--
	fmt.Printf("b = %d\n", a)

	var vi1 int16
	var vi2 int16
	var vi3 int32
	var vf1 float32 = 123.123
	var vs1 string = "hello"
	vi1, vi2, vi3, vf1, vs1 = 101, 102, 103, 123.321, "Hello go ~!"
	fmt.Println(vi1, vi2, vi3, vf1, vs1)

	var vix int16 = 1000
	fmt.Println(vi1, vi2, vix)
	vi1, vi2, vix = vi2, vix, vi1
	fmt.Println(vi1, vi2, vix)

}
