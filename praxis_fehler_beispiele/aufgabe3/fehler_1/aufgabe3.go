package aufgabe3

// CountX erwartet eine Liste von Zahlen sowie eine Zahl x.
// CountX liefert die Anzahl der Vorkommen von x in der Liste.
func CountX(list []int, x int) int {
	if len(list) == 0 {
		return 0
	}
	if list[0] == x {
		return CountX(list[1:], x) + 1
	}
	return CountX(list[1:], x)
}
