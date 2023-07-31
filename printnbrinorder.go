package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(s int) {
	var digits []int
	if s == 0 {
		z01.PrintRune('0')
	}
	a := 0
	for s > 0 {
		digits = append(digits, s%10+48)
		s = s / 10
		a++
	}
	for i := 0; i < a-1; i++ {
		for j := i + 1; j < a; j++ {
			if digits[i] > digits[j] {
				digits[i], digits[j] = digits[j], digits[i]
			}
		}
	}
	for i := 0; i < a; i++ {
		z01.PrintRune(rune(digits[i]))
	}
}
