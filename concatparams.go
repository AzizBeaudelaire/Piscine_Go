package piscine

func ConcatParams(args []string) string {
	result := ""
	for i := range args {
		if i == len(args)-1 {
			result += args[i]
		} else {
			result += args[i] + "\n"
		}
	}
	return result
}
