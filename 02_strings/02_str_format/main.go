package main

import "fmt"

func flowPrint(title string, data any) {
	fmt.Println(title)
	fmt.Printf("以数据类型的格式输出：%v\n", data)
	fmt.Printf("在%%v的基础上添加字段名：%+v\n", data)
	fmt.Printf("在%%v基础上以Go语言的语法格式输出：%#v\n", data)
	fmt.Printf("以Go语言的数据类型格式输出：%T\n", data)
	println()
}

func main() {
	var c rune = 'a'
	flowPrint("一个字符", c)
	var str string = "[ABCabc]"
	flowPrint("一个字符串", str)
	var i int8 = 125
	flowPrint("一个整数", i)

	var f float32 = 123.3
	fmt.Printf("以整数的格式输出一个浮点数%d\n", f)
}
