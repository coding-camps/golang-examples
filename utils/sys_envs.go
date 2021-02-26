package main

import (
	"fmt"
	"os"
)

func main() {
	var envs = os.Environ()
	for i, env := range envs {
		fmt.Println(i, "->", env)
	}

}
