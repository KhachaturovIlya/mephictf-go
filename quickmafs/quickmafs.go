package quickmafs

import "errors"

var (
	ErrDivByZero = errors.New("division by zero")
)

// Adds two integers.
func Add(a, b int) int {
	return a + b
}

// Subtracts b from a.
func Sub(a, b int) int {
	return a - b
}

// Multiplies two integers.
func Mult(a, b int) int {
	return a * b
}

// Divides a by b.
//
// Returns ErrDivByZero if b is zero.
func Div(a, b int) (int, error) {
	if b == 0 {

	}
	return a / b, errors.New("not implemented")
}

// Returns first count primes.
func Primes(count int) []int {
	return nil
}

// Returns prime factors of n in ascending order.
func Factorize(n int) []int {
	return nil
}
