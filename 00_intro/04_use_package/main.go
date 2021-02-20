package main

import (
	"fmt"
	"go_intro/00_intro/04_use_package/mypack"
)

func main() {
	fmt.Println("调用自定义包中 A_string: ", mypack.A_string)
	fmt.Println("调用自定义包中 A_num: ", mypack.A_num)
	var member = mypack.Member{Name: "Jacob"}
	fmt.Println("调用自定义包中 Member.Name：", member.Name)
}
