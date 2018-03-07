package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello world")
	os.Exit(0)
}

func Sum(a, b int) int {
	return a + b
}

func Substract(a, b int) int {
	return a - b
}
