package main

import (
	"fmt"
	"strconv"
)

func main() {
	var pslice []*string
	fmt.Printf("切片指针的长度: %v\n", len(pslice))
	fmt.Printf("切片指针的元素: %v\n", pslice)
	fmt.Printf("切片指针内存地址: %v\n", &pslice)

	println()
	var a, b string
	fmt.Printf("赋值前变量的地址：%v, %v\n", &a, &b)
	a, b = "a", "b"
	fmt.Printf("赋值后变量的地址：%v, %v\n", &a, &b)

	pslice = append(pslice, &a)
	pslice = append(pslice, &b)

	println()
	fmt.Printf("切片指针的长度: %v\n", len(pslice))
	fmt.Printf("切片指针的元素: %v\n", pslice)
	fmt.Printf("切片指针内存地址: %v\n", &pslice)

	// 一个指针的坑
	println()
	var ps []*string
	var s string
	for i:=1; i<=5; i++ {
		//var s string
		s = strconv.Itoa(i)
		ps = append(ps, &s)
	}
	fmt.Printf("切片指针的元素：%v\n", ps)

	//指针的指针
	println()
	var i = 123
	fmt.Printf("整型元素：%v\n", i)
	var iptr *int = &i
	fmt.Printf("整型元素指针：%v， 内容：%v\n", iptr, *iptr)
	var ipptr **int = &iptr
	fmt.Printf("整型元素指针的指针：%v, 内容：%v\n", ipptr, *ipptr)
	var ippptr ***int = &ipptr
	fmt.Printf("整型元素指针的指针的指针：%v, 内容：%v\n", ippptr, *ippptr)

	// 数组的指针
	println()
	var aptr *[3]int
	fmt.Printf("数组的指针初始化值：%v\n", aptr)
	var arr [3]int = [...]int{100,200,300}
	fmt.Printf("数组的内容：%v\n", arr)
	aptr = &arr
	fmt.Printf("数组的指针类型：%T，内容：%v\n", aptr, aptr)

}
