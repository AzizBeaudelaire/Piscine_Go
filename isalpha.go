package piscine

func IsAlpha(s string) bool {
	str := []rune(s)
	bool := false

	for i := range s {
		if str[i] >= 'A' && str[i] <= 'Z' || str[i] >= 'a' && str[i] <= 'z' || str[i] >= '0' && str[i] <= '9' {
			bool = true
		} else {
			return (false)
		}
	}
	return (bool)
}
