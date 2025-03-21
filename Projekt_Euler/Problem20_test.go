package Projekt_Euler

import "fmt"

// n! means n * (n - 1) * ... * 3 * 2 * 1.
// For example, 10! = 10 * 9 * ... * 3 * 2 * 1 = 3628800,
// and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.
// Find the sum of the digits in the number 100!.

func ExampleFactorial() {
	fmt.Println(Factorial(10))
	fmt.Println(DigitSum(Factorial(10)))
	fmt.Println(Factorial(100))
	fmt.Println(DigitSum(Factorial(100)))

	// Output:
	// 3628800
	// 27
}

// Funktion ist korrekt, da die Zahl zu lange ist um sie als int darzustellen läst sich die aufgabe nicht lösen

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}
