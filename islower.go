package piscine

func IsLower(s string) bool {
	str := []rune(s)
	bool := false

	for i := range s {
		if str[i] >= 'a' && str[i] <= 'z' {
			bool = true
		} else {
			return (false)
		}
	}
	return (bool)
}
