package main

import (
	"fmt"
)

func LogicTrue() bool {
	fmt.Print("T\t")
	return true
}

func LogicFalse() bool {
	fmt.Print("F\t")
	return false
}

func Shift1() int {
	fmt.Print("1\t")
	return 1
}

func Shift0() int {
	fmt.Print("0\t")
	return 0
}

func main() {
	fmt.Println("\n&&-1")
	fmt.Println("\nT && T ->", LogicTrue() && LogicTrue())
	fmt.Println("\n&&-2")
	fmt.Println("\nT && F ->", LogicTrue() && LogicFalse())
	fmt.Println("\n&&-3")
	fmt.Println("\nF || T ->", LogicFalse() && LogicTrue())
	fmt.Println("\n&&-4")
	fmt.Println("\nF && F ->", LogicFalse() && LogicFalse())

	fmt.Println("\n||-1")
	fmt.Println("\nT || T ->", LogicTrue() || LogicTrue())
	fmt.Println("\n||-2")
	fmt.Println("\nT || F ->", LogicTrue() || LogicFalse())
	fmt.Println("\n||-3")
	fmt.Println("\nF || T ->", LogicFalse() || LogicTrue())
	fmt.Println("\n||-4")
	fmt.Println("\nF || F ->", LogicFalse() || LogicFalse())

	fmt.Println("\n&-1")
	fmt.Println("\nT & T ->", Shift1()&Shift1())
	fmt.Println("\n&-2")
	fmt.Println("\nT & F ->", Shift1()&Shift0())
	fmt.Println("\n&-3")
	fmt.Println("\nF || T ->", Shift0()&Shift1())
	fmt.Println("\n&-4")
	fmt.Println("\nF & F ->", Shift0()&Shift0())

	fmt.Println("\n|-1")
	fmt.Println("\nT | T ->", Shift1()|Shift1())
	fmt.Println("\n|-2")
	fmt.Println("\nT | F ->", Shift1()|Shift0())
	fmt.Println("\n|-3")
	fmt.Println("\nF | T ->", Shift0()|Shift1())
	fmt.Println("\n|-4")
	fmt.Println("\nF | F ->", Shift0()|Shift0())
}
