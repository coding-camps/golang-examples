package main

import (
	"fmt"
)

func main() {

	//var num int
	//fmt.Scanf("%d", &num)
	//fmt.Println("your input number is:", num)

	// 1
	n1 := 2
	switch n1 {
	case 1:
		fmt.Println("switch-1-1")
	case 2:
		fmt.Println("switch-1-2")
	case 3:
		fmt.Println("switch-1-3")
	case 4, 5, 6:
		fmt.Println("switch-1-4,5,6")
	default:
		fmt.Println("switch-1-default")
	}

	// 2
	switch n2 := 2; n2 {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 0, 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println("其他类型")
	}

	// 3
	n3 := 2
	switch {
	case n3 < 0:
		fmt.Println("switch-3-<0")
	case n3 > 0:
		fmt.Println("switch-3->0")
	default:
		fmt.Println("switch-3-=0")
	}

	//4
	n4 := 6
	switch {
	case n4%21 == 0:
		fmt.Println("divided by 2")
		//fallthrough
	case n4%31 == 0:
		fmt.Println("divided by 3")
		//fallthrough
	default:
		fmt.Println("divided by others")
		//fallthrough
	case n4%6 == 0:
		fmt.Println("divided by 6")
	}

	// 5
	var v1 interface{} = 123.123
	switch v1.(type) {
	case int:
		fmt.Println("整数")
	case float32:
		fmt.Println("32位浮点数")
	case float64:
		fmt.Println("64位浮点数")
	default:
		fmt.Println("其他类型")
	}

	// 6
	var only int = 100
	switch only {
	default:
		fmt.Println("Only one default clause.")
	}
	//7
	switch only {
	case 100:
		fmt.Println("Only one case clause.")
	}
	switch only {
	// No case-clause or default-clause
	}
	switch v := 1; v {
	}
	switch {
	}

}
