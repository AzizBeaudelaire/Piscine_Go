package piscine

func Index(s string, toFind string) int {
	for i := 0; i < len(s); i++ {
		if len(s) <= len(toFind)+1 {
			break
		}
		if s[i:len(toFind)+i] == toFind {
			return i
		}
	}
	return (-1)
}
