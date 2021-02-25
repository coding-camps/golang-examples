package main

import "fmt"

func main() {

	// 定义
	ss := make([]int, 3)
	fmt.Printf("%v\n", ss)

	// 增加
	ss = append(ss, 10)
	fmt.Printf("%v\n", ss)

	// 复制
	var s2 = make([]int, 2)
	ccnt := copy(s2, ss)
	fmt.Printf("%v, %v, %v\n", ss, s2, ccnt)

}
