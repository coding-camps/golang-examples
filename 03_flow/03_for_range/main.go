package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Go语言"
	var ss []string = strings.Split(str, "")
	for i, s := range ss {
		fmt.Println(i, "-", s)
	}
}
