package piscine

func MakeRange(min, max int) []int {
	if min < max {
		slice := make([]int, max-min)
		for i := min; i <= max-1; i++ {
			slice[i-min] = i
		}
		return (slice)
	}
	return (nil)
}
