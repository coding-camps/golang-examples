package main

import "fmt"

func main() {

	var vi1 = 100
	fmt.Println("var int value = ", vi1)
	var ptr1 *int
	fmt.Println("ptr int default = ", ptr1)
	ptr1 = &vi1
	fmt.Println("ptr int int = ", ptr1)
	fmt.Println("ptr int value =", *ptr1)
	println()

	var vi2 = 1000
	fmt.Println("var int value = ", vi2)
	var ptr2 *int
	fmt.Println("ptr int default = ", ptr2)
	ptr2 = new(int)
	fmt.Println("ptr int init = ", ptr2)
	fmt.Println("ptr int value =", *ptr2)
	ptr2 = &vi2
	fmt.Println("ptr int def = ", ptr2)
	fmt.Println("ptr int value =", *ptr2)

	println()

	var vi3 = 10000
	fmt.Println("var int value = ", vi3)
	var ptr3 *int = new(int)
	fmt.Println("ptr int default+init = ", ptr3)
	ptr3 = &vi3
	fmt.Println("ptr int def = ", ptr3)
	fmt.Println("ptr int value =", *ptr3)


}
