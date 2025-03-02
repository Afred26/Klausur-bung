package Projekt_Euler

import "fmt"

//By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
//What is the 10001st prime number?

func ExampleNthPrime() {
	fmt.Println(NthPrime(9999))
	fmt.Println(NthPrime(10000))
	fmt.Println(NthPrime(10001))
	fmt.Println(NthPrime(10002))
	fmt.Println(NthPrime(10003))
	//Output:
	//104743
}

func NthPrime(n int) int {
	Prims := Primsbelow(200000)
	return Prims[n]
}
