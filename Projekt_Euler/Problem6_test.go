package Projekt_Euler

import (
	"fmt"
)

// The sum of the squares of the first ten natural numbers is:
// 1^2 + 2^2 + ... + 10^2 = 385.
// The square of the sum of the first ten natural numbers is:
// (1 + 2 + ... + 10)^2 = 55^2 = 3025.
// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 - 385 = 2640.
// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

func ExampleDifference() {
	fmt.Println(Difference(10))
	fmt.Println(SumSquare(10))
	fmt.Println(Pow2(Sumdown(10)))
	fmt.Println(Difference(100))
	//Output:
	//2640
	//385
	//3025
	//25164150
}

func SumSquare(n int) int {
	if n <= 0 {
		return 0
	}
	return SumSquare(n-1) + Pow2(n)
}

func Sumdown(n int) int {
	if n <= 0 {
		return 0
	}
	return Sumdown(n-1) + n
}

func Difference(n int) int {
	return Pow2(Sumdown(n)) - SumSquare(n)
}

func Pow2(n int) int {
	return n * n
}
