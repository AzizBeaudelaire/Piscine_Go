package piscine

func Swap(a *int, b *int) {
	var swap int = 0
	swap = *b
	*b = *a
	*a = swap
}
