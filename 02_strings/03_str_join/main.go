package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	m := "Hello"
	n := "Go world"

	// #1
	s1 := m + " " + n
	fmt.Println(s1)
	// #2
	s2 := fmt.Sprintf("%s %s", m, n)
	fmt.Println(s2)
	// #3
	s3 := strings.Join([]string{m, n}, ", ")
	fmt.Println(s3)
	// #4 官方推荐的拼接字符串方式
	var sbuilder strings.Builder
	sbuilder.WriteString(m)
	sbuilder.WriteString(", ")
	sbuilder.WriteString(n)
	s4 := sbuilder.String()
	fmt.Println(s4)
	// #5
	var sbuffer bytes.Buffer
	sbuffer.WriteString(m)
	sbuffer.WriteString(", ")
	sbuffer.WriteString(n)
	s5 := sbuffer.String()
	fmt.Println(s5)
}
