package aufgabe6

// Duplicates erwartet eine int-Liste list.
// Die Funktion liefert eine int-Liste mit den Elementen,
// die mehr als einmal in l1 vorkommen.
// Im Ergebnis kommt jedes Element nur einmal vor.
// Die Elemente stehen im Ergebnis in der Reihenfolge ihres ersten Auftretens.
func Duplicates(list []int) []int {
	result := []int{}
	repeats := []int{}
	remember := []int{}
	for _, el := range list {
		if !Contains(repeats, el) {
			if Contains(remember, el) {
				repeats = append(repeats, el)
			}
			if !Contains(remember, el) {
				remember = append(remember, el)
			}
		}

	}
	for _, el := range list {
		if !Contains(result, el) {
			if Contains(repeats, el) {
				result = append(result, el)
			}
		}
	}
	return result
}

// überpüfen ob el nochmal in liste ist fehlt mir
// Hilfsfunktion Contains
func Contains(l []int, x int) bool {
	result := false
	for _, el := range l {
		if el == x {
			result = true
		}
	}
	return result
}
