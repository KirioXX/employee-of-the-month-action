package main

import (
	"fmt"

	calculator "github.com/thatisuday/github-actions-golang-module"
)

func main() {
	fmt.Println("Add(6, 3) =>", calculator.Add(6, 3))
	fmt.Println("Subtract(6, 3) =>", calculator.Subtract(6, 3))
	fmt.Println("Multiply(6, 3) =>", calculator.Multiply(6, 3))
	fmt.Println("Divide(6, 3) =>", calculator.Divide(6, 3))
	fmt.Println("Fib(9) =>", calculator.Fib(9))
}
