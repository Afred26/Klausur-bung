package Projekt_Euler

import "fmt"

//The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//Find the sum of all the primes below two million.

func ExampleSumPrim() {
	fmt.Println(SumPrim(10))
	fmt.Println(SumPrim(2000000))
	fmt.Println(SumList([]int{1, 2, 3, 4}))

	// Output:
	// 17
	// 142913828922
	// 10

}

func SumPrim(n int) int {
	x := Primsbelow(n)
	return SumList(x)
}

func SumList(list []int) int {
	if len(list) == 0 {
		return 0
	}
	return list[0] + SumList(list[1:])
}
