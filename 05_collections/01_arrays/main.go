package main

import (
	"fmt"
)

func main() {

	// 一维数组定义
	println("一维数组")
	var ar11 [10]int
	fmt.Printf("数组11的类型：%T,\t 默认值：%v\n", ar11, ar11)

	for i := 5; i < 9; i++ {
		ar11[i] = i
	}
	fmt.Printf("数组11的类型：%T,\t 内容：%v\n", ar11, ar11)

	for i, _ := range ar11 {
		ar11[i] = i * 10
	}
	fmt.Printf("数组11的类型：%T,\t 内容：%v\n", ar11, ar11)

	// 定义空数组
	var ar12 []int
	fmt.Printf("数组12的类型：%T,\t 默认值：%v\n", ar12, ar12)

	//定义时全部初始化
	println()
	ar21 := [3]int{101, 102, 103}
	fmt.Printf("数组21的类型：%T,\t 内容：%v,\t长度：%d\n", ar21, ar21, len(ar21))

	var ar22 [3]int = [3]int{201, 202, 203}
	fmt.Printf("数组22的类型：%T,\t 内容：%v,\t长度：%d\n", ar22, ar22, len(ar22))

	var ar23 []int = []int{301, 302, 303}
	fmt.Printf("数组23的类型：%T,\t 内容：%v,\t长度：%d\n", ar23, ar23, len(ar23))

	var ar24 = []int{401, 402, 403}
	fmt.Printf("数组24的类型：%T,\t 内容：%v,\t长度：%d\n", ar24, ar24, len(ar24))

	var ar25 = [...]int{501, 502, 503}
	fmt.Printf("数组25的类型：%T,\t 内容：%v,\t长度：%d\n", ar25, ar25, len(ar25))

	// 定义时局部初始化
	println()
	ar31 := [5]int{1: 101, 3: 103}
	fmt.Printf("数组31的类型：%T,\t 内容：%v\t长度：%d\n", ar31, ar31, len(ar31))

	var ar32 [5]int = [5]int{1: 201, 3: 203}
	fmt.Printf("数组32的类型：%T,\t 内容：%v\t长度：%d\n", ar32, ar32, len(ar32))

	var ar33 []int = []int{1: 301, 3: 303}
	fmt.Printf("数组33的类型：%T,\t 内容：%v,\t长度：%d\n", ar33, ar33, len(ar33))

	var ar34 = []int{1: 401, 3: 403}
	fmt.Printf("数组34的类型：%T,\t 内容：%v,\t长度：%d\n", ar34, ar34, len(ar34))

	var ar35 = [...]int{1: 501, 3: 503}
	fmt.Printf("数组35的类型：%T,\t 内容：%v,\t长度：%d\n", ar35, ar35, len(ar35))

	// 二维数组定义
	println("\n二维数组")
	var ax11 [2][3]uint
	fmt.Printf("数组11的类型：%T,\t 默认值：%v\n", ax11, ax11)

	for i := 0; i < len(ax11); i++ {
		for j, _ := range ax11[i] {
			ax11[i][j] = uint((i+1)*10 + (j + 1))
		}
	}
	fmt.Printf("数组11的类型：%T,\t 内容：%v\n", ax11, ax11)
	// 定义空数组
	var ax12 []uint
	fmt.Printf("数组12的类型：%T,\t 默认值：%v\n", ax12, ax12)

	//定义时全部初始化
	println()
	ax21 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Printf("数组21的类型：%T,\t 内容：%v,\t长度：%v\n", ax21, ax21, [2]int{len(ax21), len(ax21[0])})

	var ax22 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Printf("数组22的类型：%T,\t 内容：%v,\t长度：%v\n", ax22, ax22, [2]int{len(ax22), len(ax22[0])})

	var ax23 [][]int = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Printf("数组23的类型：%T,\t 内容：%v,\t长度：%v\n", ax23, ax23, [2]int{len(ax23), len(ax23[0])})

	var ax24 = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Printf("数组24的类型：%T,\t 内容：%v,\t长度：%v\n", ax24, ax24, [2]int{len(ax24), len(ax24[0])})

	var ax25 = [...][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Printf("数组25的类型：%T,\t 内容：%v,\t长度：%v\n", ax25, ax25, [2]int{len(ax25), len(ax25[0])})

	// 数组操作
}
