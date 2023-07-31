package piscine

func TrimAtoi(s string) int {
	isnegative := false
	asnumber := false
	result := ""
	for i := range s {
		if s[i] >= '0' && s[i] <= '9' {
			asnumber = true
			result += string(s[i])
		}
		if s[i] == '-' && asnumber == false {
			isnegative = true
		}
	}
	if isnegative == true {
		return -Atoi(result)
	}
	return Atoi(result)
}
