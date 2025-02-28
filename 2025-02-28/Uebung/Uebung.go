package uebung

/*
func types(i int) {
	fmt.Printf("%T", i)
}
*/

func cross_sum_weight(n, c int) int {
	if n == 0 {
		return 0
	}
	return cross_sum_weight(n/10, c+1) + c*(n%10)
}

func cross_sum(n int) int {
	if n == 0 {
		return 0
	}
	return cross_sum(n/10) + (n % 10)
}
