package Projekt_Euler

import (
	"fmt"
	"slices"
)

// A palindromic number reads the same both ways.
// The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 * 99.
// Find the largest palindrome made from the product of two 3-digit numbers.
func Example_palindrom() {

	fmt.Println(IsPalendrom(101))
	fmt.Println(IsPalendrom(100000010))
	x := GenPalindromBetween(10100, 997002)
	fmt.Println(x)
	fmt.Println(Greatest(x))

	//Output:
	// true
	// false
	//996699
}

// Bisher generiere ich alle Palindrome von 100*101 bis 998*999.
// Ich filtere aber nicht nach den Zahlen mit den richtigen Teilern.

func IsPalendrom(n int) bool {
	numf := NumSlice(n)
	numb := NumSlice(n)
	slices.Reverse(numb)
	return slices.Equal(numf, numb)

}

func NumSlice(n int) []int {
	num := []int{}
	for n > 0 {
		num = append(num, n%10)
		n /= 10
	}
	slices.Reverse(num)
	return num
}

func GenPalindromBetween(min, max int) []int {
	if min == max {
		return []int{}
	}
	if IsPalendrom(min) {
		temp := []int{min}
		return append(temp, GenPalindromBetween(min+1, max)...)
	}

	return GenPalindromBetween(min+1, max)
}
