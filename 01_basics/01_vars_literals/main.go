package main

import (
	"fmt"
)

func main() {

	// variables
	var a1 int
	var a2, a3 int
	var a4 int = 40
	var a5, a6 int = 50, 60
	a7 := 70
	a8, a9 := 80, 90
	var (
		x1     int
		x2, x3 int
		x4     int = 40
		x5, x6 int = 50, 60
		x7         = 70
		x8, x9     = 80, 90
	)

	a1 = 10
	a2, a3 = 20, 30
	// 10 20 30 40 50 60 70 80 90
	fmt.Println(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	x1 = 10
	x2, x3 = 20, 30
	// 10 20 30 40 50 60 70 80 90
	fmt.Println(x1, x2, x3, x4, x5, x6, x7, x8, x9)

	// constants
	const b1 int = 10
	const b2, b3 int = 20, 30
	const b4 = 40
	const b5, b6 = 50, 60
	const (
		y1     int = 10
		y2, y3 int = 20, 30
		y4         = 40
		y5, y6     = 50, 60
	)
	// 10 20 30 40 50 60
	fmt.Println(b1, b2, b3, b4, b5, b6)
	// 10 20 30 40 50 60
	fmt.Println(y1, y2, y3, y4, y5, y6)

	// iota intro
	const c1 = iota
	const c2 = iota
	const (
		c3 = "[0]"
		c4 = iota
		c5
		c6 = "[ok]"
		c7
		c8 = iota
		c9
		_
		c10
	)
	// 0 0 [0] 1 2 [ok] [ok] 5 6 8
	fmt.Println(c1, c2, c3, c4, c5, c6, c7, c8, c9, c10)

	// a good example of iota
	const (
		_          = iota
		KB float64 = 1 << (10 * iota)
		MB
		GB
		TB
	)
	fmt.Printf("K -> KB : %.0f\n", KB)
	fmt.Printf("K -> MB : %.0f\n", MB)
	fmt.Printf("K -> GB : %.0f\n", GB)
	fmt.Printf("K -> TB : %.0f\n", TB)
}
