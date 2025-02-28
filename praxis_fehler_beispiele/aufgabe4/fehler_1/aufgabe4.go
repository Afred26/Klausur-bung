package aufgabe4

// ElementProducts erwartet zwei int-Listen l1 und l2.
// Sie liefert eine Liste, die an jeder Position
// jeweils das Produkt der beiden Elemente enthält.
// Falls eine Position nur in einer Liste vorkommt,
// soll dieses Element ins Ergebnis übernommen werden.
func ElementProducts(l1, l2 []int) []int {
	if len(l1) == 0 {
		return l2
	}
	if len(l2) == 0 {
		return l1
	}
	result := []int{l1[0] * l2[0]}
	return append(result, ElementProducts(l1[1:], l2[1:])...)
}
