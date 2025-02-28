package uebung

import "fmt"

func Example_types() {
	x := 1
	types(x)

	//Output:
	// int
}

func Example_int_float_type() {
	x1 := 15
	x2 := 2.3

	fmt.Printf("x1: %T\n", x1)
	fmt.Printf("x2: %T\n", x2)

	//Output:
	//x1: int
	//x2: float64
}
