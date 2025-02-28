package Projekt_Euler

import "fmt"

//2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
//What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

func ExampleDivisors() {
	x := 200

	//fmt.Println(Divisors(x))
	y := PrimFaktorsBelow(x)
	for i, el := range y {
		fmt.Printf("Die Primfaktoren von %d sind: %v\n", i+2, el)
	}

	//Output:
}

func Divisors(n int) []PrimFaktor {
	Prims := Primsbelow(n + 1)
	result := []PrimFaktor{}
	for _, el := range Prims {
		num := MultDiv(n, el)
		if num.Potenz != 0 {
			result = append(result, num)
		}
	}
	return result
}

type PrimFaktor struct {
	Primzahl int
	Potenz   int
}

func MultDiv(Zahl, Primzahl int) PrimFaktor {
	if Zahl%Primzahl != 0 {
		return PrimFaktor{0, 0}
	}
	result := PrimFaktor{Primzahl: Primzahl, Potenz: 0}
	for Zahl%Primzahl == 0 {
		result.Potenz++
		Zahl /= Primzahl
	}
	return result
}

func PrimFaktorsBelow(n int) [][]PrimFaktor {
	result := make([][]PrimFaktor, n-1)

	for i := 2; i < n+1; i++ {
		result[i-2] = Divisors(i)
	}
	return result
}
