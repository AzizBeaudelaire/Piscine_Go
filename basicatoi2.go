package piscine

func BasicAtoi2(s string) int {
	x := 0
	for _, num := range s {
		if num > 47 && num < 58 {
			x = x*10 + int(num-48)
		} else {
			return (0)
		}
	}
	return (x)
}
