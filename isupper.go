package piscine

func IsUpper(s string) bool {
	str := []rune(s)
	bool := false

	for i := range s {
		if str[i] >= 'A' && str[i] <= 'Z' {
			bool = true
		} else {
			return (false)
		}
	}
	return (bool)
}
