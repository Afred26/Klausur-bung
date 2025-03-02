package Projekt_Euler

import "fmt"

//A Pythagorean triplet is a set of three natural numbers, a < b < c, for which, a^2 + b^2 = c^2.
//For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
//There exists exactly one Pythagorean triplet for which a + b + c = 1000.
//Find the product abc.

func Example() {
	x1 := Pythagoreantriplet{a: 3, b: 4, c: 5}
	x2 := Pythagoreantriplet{a: 3, b: 4, c: 993}
	fmt.Println(x1.IsPythagorean())
	fmt.Println(x2.IsPythagorean())
	fmt.Println(x2.SumThousen())
	fmt.Println(x1.Product())
	fmt.Println(GetIntList(10))
	result := FindTriplets(1000)
	fmt.Println(result[1].Product())

	//Output:
}

type Pythagoreantriplet struct {
	a int
	b int
	c int
}

func (p Pythagoreantriplet) IsPythagorean() bool {
	if Pow2(p.a)+Pow2(p.b) == Pow2(p.c) {
		return true
	} else {
		return false
	}
}

func (p Pythagoreantriplet) SumThousen() bool {
	if p.a+p.b+p.c == 1000 {
		return true
	} else {
		return false
	}
}

func (p Pythagoreantriplet) Product() int {
	return p.a * p.b * p.c
}

func GetIntList(n int) []int {
	if n < 0 {
		return []int{}
	}
	return append(GetIntList(n-1), n)
}

func FindTriplets(n int) []Pythagoreantriplet {
	list := GetIntList(n)
	p := Pythagoreantriplet{a: 0, b: 0, c: 0}
	result := []Pythagoreantriplet{}

	for _, ela := range list {
		p.a = ela
		for _, elb := range list {
			p.b = elb
			if p.a+p.b < n {
				p.c = n - (p.a + p.b)
				if p.IsPythagorean() {
					result = append(result, p)
				}
			}
		}
	}

	return result
}
