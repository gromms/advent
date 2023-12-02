package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	fileSystem := os.DirFS("./")
	file, error := fs.ReadFile(fileSystem, "input")

	if error != nil {
		fmt.Println(error)
	}

	firstDigit := -1
	lastDigit := -1

	sum := 0

	for _, char := range file {
		if char >= 48 && char <= 57 {
			num := int(char) - 48
			if firstDigit == -1 {
				firstDigit = num
			}
			lastDigit = num
		} else if char == 10 {
			num := firstDigit*10 + lastDigit
			//         fmt.Printf("Line digits: %d - %d; Number: %d\n", firstDigit, lastDigit, num)
			sum += num
			firstDigit = -1
			lastDigit = -1
		}

	}

	fmt.Println(sum)
}
