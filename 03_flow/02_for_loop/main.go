package main

import "fmt"

func main() {
	var str string = "golang"
	var i int = 0

	// 1
	fmt.Printf("\n%d -> ", 1)
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c  ", str[i])
	}
	// 2
	fmt.Printf("\n%d -> ", 2)
	i = 0
	for ; i < len(str); i++ {
		fmt.Printf("%c  ", str[i])
	}
	// 3
	fmt.Printf("\n%d -> ", 3)
	for j := 0; ; j++ {
		if j >= len(str) {
			break
		}
		fmt.Printf("%c  ", str[j])
	}
	// 4
	fmt.Printf("\n%d -> ", 4)
	for j := 0; j < len(str); {
		fmt.Printf("%c  ", str[j])
		j++
	}

	// 5
	fmt.Printf("\n%d -> ", 5)
	i = 0
	for ; ; i++ {
		if i >= len(str) {
			break
		}
		fmt.Printf("%c  ", str[i])
	}

	// 6
	fmt.Printf("\n%d -> ", 6)
	i = 0
	for i < len(str) {
		fmt.Printf("%c  ", str[i])
		i++
	}

	//7
	fmt.Printf("\n%d -> ", 7)
	i = 0
	// for ;i < len(str); {  注意这里两个分号，IDE一般会自动删除他们
	for ;i < len(str); {
		fmt.Printf("%c  ", str[i])
		i++
	}

	// 8
	fmt.Printf("\n%d -> ", 8)
	for j := 0; ; {
		if j >= len(str) {
			break
		}
		fmt.Printf("%c  ", str[j])
		j++
	}

	// 9
	fmt.Printf("\n%d -> ", 9)
	i = 0
	for {
		if i >= len(str) {
			break
		}
		fmt.Printf("%c  ", str[i])
		i++
	}

	// 10
	fmt.Printf("\n%d -> ", 10)
	i = 0
	// for ;;{ 注意这里有两个分号，一般情况IDE会自动删除。
	for ;;{
		if i >= len(str) {
			break
		}
		fmt.Printf("%c  ", str[i])
		i++
	}
}
