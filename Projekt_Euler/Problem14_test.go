package Projekt_Euler

import "fmt"

//The following iterative sequence is defined for the set of positive integers:
// n -> n/2 (n is even)
// n -> 3n + 1 (n is odd)
//Using the rule above and starting with $13$, we generate the following sequence:
// 13 -> 40 -> 20 -> 10 -> 5 -> 16 -> 8 -> 4 -> 2 -> 1.
//It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.
//Which starting number, under one million, produces the longest chain?
//Once the chain starts the terms are allowed to go above one million.

func ExampleCollatz() {
	biggest := []int{0, 0}
	for i := 1; i < 1000000; i++ {
		x := Collatz(i, 1)
		if x > biggest[0] {
			biggest[0] = x
			biggest[1] = i
		}
	}
	fmt.Println(biggest)

	// Output:

}

func Collatz(n, l int) int {
	if n == 1 {
		return l
	} else if n%2 == 0 {
		return Collatz(n/2, l+1)
	} else {
		return Collatz((3*n)+1, l+1)
	}
}
