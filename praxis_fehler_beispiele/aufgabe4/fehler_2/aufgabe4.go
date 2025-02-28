package aufgabe4

// ElementProducts erwartet zwei int-Listen l1 und l2.
// Sie liefert eine Liste, die an jeder Position
// jeweils das Produkt der beiden Elemente enthält.
// Falls eine Position nur in einer Liste vorkommt,
// soll dieses Element ins Ergebnis übernommen werden.
func ElementProducts(l1, l2 []int) []int {
	min := min(len(l1), len(l2))
	result := []int{}
	for i := 0; i < min; i++ {
		result = append(result, l1[i]*l2[i])
	}
	result = append(result, l1[min:]...)
	result = append(result, l2[min:]...)
	return result
}
