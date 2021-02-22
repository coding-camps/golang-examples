package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	m := "Hello,golang."
	n := "你好，Go语言。"

	fmt.Println(len(m), len(n))
	fmt.Println(utf8.RuneCountInString(m), utf8.RuneCountInString(n))
	fmt.Println(len([]rune(m)), len([]rune(n)))

	println(n)
	for i := 0; i < len(n); i++ {
		fmt.Printf("[%c]--%d\n", n[i], n[i])
	}

	println()
	ns := []rune(n)
	for i := 0; i < len(ns); i++ {
		fmt.Printf("[%c]--%d\n", ns[i], ns[i])
	}

	println()
	for i, s := range n {
		fmt.Printf("%-2d: [%c]--%d\n", i, s, s)
	}

}
