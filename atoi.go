package piscine

func Atoi(s string) int {
	x := 0
	for _, i := range s {
		x = x*10 + int(i-48)
		if s[0] == '+' {
			return BasicAtoi2(s[1:])
		}
		if s[0] == '-' {
			return -BasicAtoi2(s[1:])
		}
		if i < 48 || i > 57 {
			return 0
		}
	}
	return x
}
