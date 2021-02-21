package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i int = 128
	var f = 123.501
	var f1 float32 = 123.511
	fmt.Println(i, f, f1)
	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(f), reflect.TypeOf(f1))
	f = float64(f1)
	fmt.Println(f)

	var c1 = complex(1, 2)
	var c2 complex64 = complex(11, 22)
	fmt.Println(c1, c2)
	fmt.Println(reflect.TypeOf(c1), reflect.TypeOf(c2))
}
