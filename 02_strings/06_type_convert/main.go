package main

import (
	"fmt"
	"strconv"
)

func main() {

	println()
	var i1 int = 100
	var i2 int8 = 108
	var i3 int64 = 1008
	var f1 float32 = 103.6
	var f2 float64 = 1003.603
	fmt.Printf("integer numbers: %d, %d, %d\n", i1, i2, i3)
	fmt.Printf("float numbers: %f, %f\n", f1, f2)

	// int -> float
	fmt.Printf("int -> float: %f, %f\n", float32(i1), float64(i2))
	// float -> int, will miss the decimal part
	fmt.Printf("float -> int: %d, %d\n", int(f1), int64(f2))

	println()
	// str -> int
	c1, err1 := strconv.Atoi("100.12")
	fmt.Printf("%T->%d\nerror msg: %v\n", c1, c1, err1)
	c1, err1 = strconv.Atoi("100")
	fmt.Printf("%T->%d\nerror msg: %v\n", c1, c1, err1)
	// str -> int8, int16, int32, int64
	c2, _ := strconv.ParseInt("123", 10, 8)
	fmt.Printf("str->int8: %T, %d\n", c2, c2)
	c3, _ := strconv.ParseInt("1230", 10, 16)
	fmt.Printf("str->int8: %T, %d\n", c3, c3)
	c4, _ := strconv.ParseInt("1230", 10, 32)
	fmt.Printf("str->int8: %T, %d\n", c4, c4)
	c5, _ := strconv.ParseInt("1230", 10, 64)
	fmt.Printf("str->int8: %T, %d\n", c5, c5)
	// int -> str
	c6 := strconv.Itoa(100)
	fmt.Printf("%T->%s\n", c6, c6)

	println()
	// float -> str
	fmt.Printf("f1=%f, f2=%f\n", f1, f2)
	f1s1 := strconv.FormatFloat(float64(f1), 'f', 10, 32)
	f1s2 := strconv.FormatFloat(float64(f1), 'f', 1, 32)
	fmt.Printf("f1s1 = %s, f1s2=%s\n", f1s1, f1s2)
	f2s1 := strconv.FormatFloat(f2, 'b', 1, 32)
	f2s2 := strconv.FormatFloat(f2, 'b', 1, 64)
	fmt.Printf("f2s1 = %s, f2s2=%s\n", f2s1, f2s2)
	// str -> float
	sn1 := "123.123"
	sn2 := "3.1415926"
	n1, _ := strconv.ParseFloat(sn1, 32)
	n2, _ := strconv.ParseFloat(sn2, 64)
	fmt.Printf("n1 = %f, n2 = %f\n", n1, n2)
}
