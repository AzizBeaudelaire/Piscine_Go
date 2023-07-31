package piscine

func BasicAtoi(s string) int {
	x := 0
	for _, num := range s {
		x = x*10 + int(num-48)
	}
	return (x)
}
