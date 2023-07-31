package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	nb := n
	table := []int{}
	if n < 0 {
		z01.PrintRune('-')
	} else if n == 0 {
		z01.PrintRune(rune(48))
	}
	for nb != 0 {
		if n < 0 {
			table = append(table, -(nb%10)+48)
		} else {
			table = append(table, (nb%10)+48)
		}
		nb = nb - nb%10
		nb = nb / 10
	}
	for i := len(table) - 1; i > -1; i-- {
		z01.PrintRune(rune(table[i]))
	}
}
