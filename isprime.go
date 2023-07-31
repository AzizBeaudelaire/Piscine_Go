package piscine

func IsPrime(nb int) bool {
	if nb == 1 {
		return false
	} else if nb > 1 {
		for i := 2; i <= nb-1; i++ {
			if nb%i == 0 {
				return (false)
			}
		}
		return (true)
	}
	return (false)
}
