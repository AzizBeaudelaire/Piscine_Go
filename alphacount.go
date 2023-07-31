package piscine

func AlphaCount(s string) int {
	str := []rune(s)
	result := 0

	for n := range s {
		if str[n] >= 'A' && str[n] <= 'Z' || str[n] >= 'a' && str[n] <= 'z' {
			result++
		}
	}
	return (result)
}
