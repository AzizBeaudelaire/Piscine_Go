package piscine

func NRune(s string, n int) rune {
	cds := []rune(s)
	if n > 0 && n <= len(s) {
		return cds[n-1]
	}
	return (0)
}
