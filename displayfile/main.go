package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("quest8.txt")
	arg := os.Args
	arr := make([]byte, 14)

	if len(arg) == 1 {
		fmt.Println("File name missing")
	} else if len(arg) > 2 {
		fmt.Println("Too many arguments")
	} else {
		file.Read(arr)
		fmt.Println(string(arr))
	}
}
