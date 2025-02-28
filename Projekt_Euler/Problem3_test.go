package Projekt_Euler

import (
	"fmt"
	"slices"
)

//The prime factors of 13195 are 5, 7, 13 and 29.
//What is the largest prime factor of the number 600851475143?

func ExamplePrims() {
	//x := 600851475143
	x := 5000000
	//Prim := Prims(x)
	divisor := PrimDivisor(x)
	y := Greatest(divisor)
	//fmt.Println(Prim)
	fmt.Println(divisor)
	fmt.Println(y)

	//Output:
	//[5 2]
	//5
}

func Prims(max int) []int {
	//sive := []int{}
	prims := []int{}
	for i := 2; i < max; i++ {
		if !Div(prims, i) {
			prims = append(prims, i)
		}
	}

	return prims
}

func PrimDivisor(n int) []int {
	p := Prims(n)
	divisors := []int{}
	for _, el := range p {
		if n%el == 0 {
			divisors = append(divisors, el)
		}
	}
	return divisors
}

func Div(list []int, d int) bool {
	for _, el := range list {
		if d%el == 0 {
			return true
		}
	}
	return false
}

func Greatest(list []int) int {
	slices.Sort(list)
	slices.Reverse(list)
	return list[0]
}
