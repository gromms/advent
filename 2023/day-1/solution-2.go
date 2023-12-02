package main

import (
	"fmt"
	"io/fs"
	"os"
)

func setDigits(first *int, last *int, digit int) {
	if *first == -1 {
		*first = digit
	}
	*last = digit
}

func resetDigits(first *int, last *int) {
	*first = -1
	*last = -1
}

func main() {
	fileSystem := os.DirFS("./")
	file, error := fs.ReadFile(fileSystem, "input")

	if error != nil {
		fmt.Println(error)
	}

	firstDigit := -1
	lastDigit := -1

	sum := 0

	line := 0
	for index, char := range file {
		if char >= 48 && char <= 57 {
			num := int(char) - 48
			if firstDigit == -1 {
				firstDigit = num
			}
			lastDigit = num
		} else if index+3 < len(file) && string(file[index:index+3]) == "one" {
			setDigits(&firstDigit, &lastDigit, 1)
		} else if index+3 < len(file) && string(file[index:index+3]) == "two" {
			setDigits(&firstDigit, &lastDigit, 2)
		} else if index+5 < len(file) && string(file[index:index+5]) == "three" {
			setDigits(&firstDigit, &lastDigit, 3)
		} else if index+4 < len(file) && string(file[index:index+4]) == "four" {
			setDigits(&firstDigit, &lastDigit, 4)
		} else if index+4 < len(file) && string(file[index:index+4]) == "five" {
			setDigits(&firstDigit, &lastDigit, 5)
		} else if index+3 < len(file) && string(file[index:index+3]) == "six" {
			setDigits(&firstDigit, &lastDigit, 6)
		} else if index+5 < len(file) && string(file[index:index+5]) == "seven" {
			setDigits(&firstDigit, &lastDigit, 7)
		} else if index+5 < len(file) && string(file[index:index+5]) == "eight" {
			setDigits(&firstDigit, &lastDigit, 8)
		} else if index+4 < len(file) && string(file[index:index+4]) == "nine" {
			setDigits(&firstDigit, &lastDigit, 9)
		} else if char == 10 {
			num := firstDigit*10 + lastDigit
			//         fmt.Printf("Line digits: %d - %d; Number: %d\n", firstDigit, lastDigit, num)
			sum += num
			resetDigits(&firstDigit, &lastDigit)
			line++
		}

	}

	fmt.Println(sum)
}
