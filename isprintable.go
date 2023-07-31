package piscine

func IsPrintable(s string) bool {
	str := []rune(s)
	bool := false
	for i := range s {
		if str[i] > 31 {
			bool = true
		} else {
			return (false)
		}
	}
	return (bool)
}
