package calculator

import fib "github.com/t-pwk/go-fibonacci"

// Add returns the sum of two integers
func Add(a int, b int) int {
	return a + b
}

// Subtract returns the subtraction of two integers
func Subtract(a int, b int) int {
	return a - b
}

// Multiply returns the multiplication of two integers
func Multiply(a int, b int) int {
	return a * b
}

// Divide returns the division of two integers
func Divide(a int, b int) int {
	return a / b
}

// Fib returns the nth fibonacci number
func Fib(n uint) uint64 {
	return fib.Fibonacci(n)
}
