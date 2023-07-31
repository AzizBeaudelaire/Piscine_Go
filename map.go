package piscine

func Map(f func(int) bool, a []int) []bool {
	var array []bool
	for _, i := range a {
		array = append(array, f(i))
	}
	return (array)
}
