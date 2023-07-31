package piscine

func AppendRange(min, max int) []int {
	var slice []int
	if min >= max {
		return (nil)
	}
	for i := min; i <= max-1; i++ {
		slice = append(slice, i)
	}
	return (slice)
}
