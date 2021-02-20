package main

import "fmt"

func main() {

	var name, age string
	fmt.Printf("Please input your name:")
	fmt.Scanln(&name)
	fmt.Printf("Please input your age:")
	fmt.Scanln(&age)
	fmt.Printf("Your name is %v, age is %v.\n", name, age)

}
