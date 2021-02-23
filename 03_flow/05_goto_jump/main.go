package main

import "fmt"

func main() {
	var intList = []int{1, 2, 3, 4, 5, 6, 7, 8}

	for _, k := range intList {

		fmt.Print(k, ", ")

		if k > 5 {
			goto endFlag
		}
	}
endFlag:
}
