package piscine

func Sqrt(nb int) int {
	if nb == 1 {
		return (1)
	} else if nb > 1 {
		for i := 2; i < nb; i++ {
			if i*i == nb {
				return (i)
			} else {
				continue
			}
		}
	}
	return (0)
}
