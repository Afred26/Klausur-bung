package aufgabe2

// ExcludeBetween erwartet eine Liste und zwei Zahlen m und n.
// Die Funktion liefert eine Liste mit allen Elementen x, für die gilt: m < x < n.
func ExcludeBetween(list []int, m, n int) []int {
	if len(list) == 0 {
		return []int{}
	}
	if list[0] > m && list[0] < n {
		return append(ExcludeBetween(list[1:], m, n), list[0])
	}
	return ExcludeBetween(list[1:], m, n)
}
