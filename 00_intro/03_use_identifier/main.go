package main

import (
	"fmt"
	_ "net/http/pprof"
)

func myfunc() (int, string) {
	a := 10
	b := "hello"
	return a, b
}

func main() {
	a, _ := myfunc()
	fmt.Printf("仅使用第一个函数值 %d\n", a)
}
