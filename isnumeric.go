package piscine

func IsNumeric(s string) bool {
	str := []rune(s)
	bool := false
	for i := range s {
		if str[i] >= '0' && str[i] <= '9' {
			bool = true
		} else {
			return (false)
		}
	}
	return (bool)
}
