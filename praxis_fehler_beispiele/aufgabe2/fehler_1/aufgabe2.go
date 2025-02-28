package aufgabe2

// ExcludeBetween erwartet eine Liste und zwei Zahlen m und n.
// Die Funktion liefert eine Liste mit allen Elementen x, für die gilt: m < x < n.
func ExcludeBetween(list []int, m, n int) []int {
	result := []int{}
	for _, el := range list {
		if el > m && el < n {
			result = append(result, el)
		}
	}

	return result
}
