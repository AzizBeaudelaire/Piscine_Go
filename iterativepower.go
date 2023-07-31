package piscine

func IterativePower(nb int, power int) int {
	result := 1
	if power == 0 {
		return (1)
	} else if power > 0 {
		for i := 0; i < power; i++ {
			result = result * nb
		}
		return (result)
	}
	return (0)
}
