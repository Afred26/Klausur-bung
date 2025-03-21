package Projekt_Euler

import (
	"fmt"
)

//2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
//What is the sum of the digits of the number 2^1000?

func ExampleDigitSum() {
	fmt.Println(Pow(2, 15))
	fmt.Println(DigitSum(Pow(2, 15)))
	fmt.Println(DigitSum(Pow(2, 1000)))

	// Output:
	// 32768
	// 26
}

// Funktion ist korrekt
// Stellen von int nicht genug um Zahl darzustellen

func DigitSum(n int) int {
	list := []int{}
	for n > 0 {
		list = append(list, n%10)
		n = n / 10
	}
	return SumList(list)
}

// berechnet x^y
func Pow(x, y int) int {
	if y == 0 {
		return 1
	}
	if y < 0 {
		return 1 / Pow(x, -y)
	}
	return x * Pow(x, y-1)
}
