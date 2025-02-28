package Projekt_Euler

import "fmt"

// If we list all the natural numbers below 10
// that are multiples of 3 or 5, we get 3, 5, 6 and 9.
// The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000

func Example_projekt1() {
	fmt.Println(Sum(9))
	fmt.Println(Sum(10))
	fmt.Println(Sum(999))

	// Output:
	// 23
	// 33
	// 233168
}

func Mult_3_5(n int) bool {
	if n%3 == 0 || n%5 == 0 {
		return true
	}
	return false
}

func Sum(n int) int {
	if n == 0 {
		return 0
	}
	if Mult_3_5(n) {
		return Sum(n-1) + n
	}
	return Sum(n - 1)
}
