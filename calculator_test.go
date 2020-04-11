package calculator

import (
	"testing"
)

// Add returns the sum of two integers
func TestAdd(t *testing.T) {
	if Add(6, 3) != 9 {
		t.Fail()
	}
}

// Subtract returns the subtraction of two integers
func TestSubtract(t *testing.T) {
	if Subtract(6, 3) != 3 {
		t.Fail()
	}
}

// Multiply returns the multiplication of two integers
func TestMultiply(t *testing.T) {
	if Multiply(6, 3) != 18 {
		t.Fail()
	}
}

// Divide returns the division of two integers
func TestDivide(t *testing.T) {
	if Divide(6, 3) != 2 {
		t.Fail()
	}
}

// Fib returns the nth fibonacci number
func TestFib(t *testing.T) {
	if Fib(9) != 34 {
		t.Fail()
	}
}
