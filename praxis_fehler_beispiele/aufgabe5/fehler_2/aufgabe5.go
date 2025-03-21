package aufgabe5

// IsChain erwartet eine Liste von Dominoe-Objekten.
// Die Funktion prüft, ob diese Liste eine Kette bildet,
// die nach den Domino-Regeln erlaubt wäre.
// Bei einer solchen Kette muss immer die rechte Seite eines Steins
// gleich der linken Seite des nächsten Steins sein.
func IsChain(dominoes []Dominoe) bool {
	if len(dominoes) <= 1 {
		return true
	}
	if dominoes[0].Right == dominoes[1].Left {
		return IsChain(dominoes[1:])
	}
	return false
}

// Dominoe repräsentiert einen Domino-Stein mit zwei Zahlen.
type Dominoe struct {
	Left  int
	Right int
}
