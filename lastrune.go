package piscine

func LastRune(s string) rune {
	cds := []rune(s)
	i := 0
	for range s {
		i++
	}
	return (cds[i-1])
}
