package main

import (
	"fmt"
)

func main() {
	var intList = []int{1, 2, 3, 4, 5, 6, 7, 8}

	// 1 break
	for _, d := range intList {
		fmt.Print(" [")
		if d > 5 {
			break
		}
		fmt.Printf("%d", d)
		fmt.Print("] ")
	}
	println()

	// 2 continue
	for _, d := range intList {
		fmt.Print(" [")
		if d > 5 {
			continue
		}
		fmt.Printf("%d", d)
		fmt.Print("] ")
	}
	println()

	// print 2D array
	var intMatrix = [][]int{
		{11, 12, 13, 14, 15, 16, 17, 18},
		{21, 22, 23, 24, 25, 26, 27, 28},
		{31, 32, 33, 34, 35, 36, 37, 38},
		{41, 42, 43, 44, 45, 46, 47, 48},
		{51, 52, 53, 54, 55, 56, 57, 58},
		{61, 62, 63, 64, 65, 66, 67, 68},
		{71, 72, 73, 74, 75, 76, 77, 78},
		{81, 82, 83, 84, 85, 86, 87, 88},
	}
	println()
	for i := 0; i < len(intMatrix); i++ {
		for j := 0; j < len(intMatrix[i]); j++ {
			fmt.Print(" [")
			fmt.Printf("%d", intMatrix[i][j])
			fmt.Print("] ")
		}
		println()
	}

	// 3 continue inner & outer loop
	println()
outerLoop1:
	for i := 0; i < len(intMatrix); i++ {
	innerLoop1:
		for j := 0; j < len(intMatrix[i]); j++ {

			if i%2 == 0 {
				continue outerLoop1
			}
			if j%2 == 0 {
				continue innerLoop1
			}

			fmt.Print(" [")
			fmt.Printf("%d", intMatrix[i][j])
			fmt.Print("] ")
		}
		println()
	}

	// 3 break inner & outer loop
	println()
outerLoop2:
	for i := 0; i < len(intMatrix); i++ {
	innerLoop2:
		for j := 0; j < len(intMatrix[i]); j++ {

			if i > 3 {
				break outerLoop2
			}

			if j > 3 {
				break innerLoop2
			}

			fmt.Print(" [")
			fmt.Printf("%d", intMatrix[i][j])
			fmt.Print("] ")
		}
		println()
	}

	// for switch continue/break
	println()
outerFor:
	for _, item := range intList {
		fmt.Printf("\n%d", item)

		switch {
		case item > 4:
			// break 不可以
			break outerFor
		}

		fmt.Print(" - ")

		switch {
		case item > 2:
			// continue 可以
			continue outerFor
		}
		fmt.Print(" * ")

	}
	println()

}
