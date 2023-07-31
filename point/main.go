package main

import "github.com/01-edu/z01"

type point struct {
	x, y rune
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	for _, r := range "x = " {
		z01.PrintRune(r)
	}
	z01.PrintRune(points.x/10 + 48)
	z01.PrintRune(points.x%10 + 48)
	for _, r := range ", y = " {
		z01.PrintRune(r)
	}
	z01.PrintRune(points.y/10 + 48)
	z01.PrintRune(points.y%10 + 48)

	z01.PrintRune('\n')
}
