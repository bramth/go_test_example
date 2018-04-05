package main

import (
	"fmt"
	"os"

	"github.com/wendyadi/go_test_example/user"
)

func main() {
	fmt.Println("Hello world")
	user.Init()
	os.Exit(0)
}

func Sum(a, b int) int {
	res := a + b
	if res > 10 {
		return 100
	}
	return res
}

func Sub(a, b int) int {
	return a - b
}
