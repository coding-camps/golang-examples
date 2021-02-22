package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "你好 Golang World！"
	fmt.Println(s)
	i := strings.Index(s, "World")
	fmt.Println("index of \"World\" is:", i)
	i = strings.LastIndex(s, "World")
	fmt.Println("last index of \"World\" is:", i)
	sx := s[i:]
	fmt.Println("Slice of the string is:", sx)
	println()

	sp := strings.Split(s, " ")
	fmt.Printf("type of after spliting:%T\n", sp)
	for j, sj := range sp {
		fmt.Println(j, "->", sj)
	}
	println()
	for j := 0; j < len(sp); j++ {
		fmt.Println(j, "->", sp[j])
	}

	println()
	sr1 := strings.Replace(s, "o", "*", -1)
	sr2 := strings.Replace(s, "o", "*", 1)
	sr3 := strings.Replace(s, "o", "*", 0)
	sr4 := strings.Replace(s, "o", "*", -2)
	sr5 := strings.Replace(s, "o", "*", 100)
	fmt.Println(sr1)
	fmt.Println(sr2)
	fmt.Println(sr3)
	fmt.Println(sr4)
	fmt.Println(sr5)
}
