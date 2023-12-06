package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseNumbers(line string, sep string) int {

	var nums []string
	a := strings.Split(line, sep)[1]
	for _, s := range strings.Split(a, " ") {
		if s != "" {
			nums = append(nums, s)
		}
	}
	result, _ := strconv.Atoi(strings.Join(nums, ""))
	return result
}

func main() {

	fileName := os.Args[1]

	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)

	scanner.Scan()

	duration := parseNumbers(scanner.Text(), "Time:")
	fmt.Printf("%d\n", duration)

	scanner.Scan()

	distance := parseNumbers(scanner.Text(), "Distance:")
	fmt.Printf("%d\n", distance)

	minDistance := distance + 1

	halfSqrt := math.Sqrt(float64(duration*duration-4*minDistance)) / 2
	halfDuration := float64(duration) / 2

	min := int(math.Ceil(halfDuration - halfSqrt))
	max := int(math.Floor(halfDuration + halfSqrt))

	winningSolutions := max - min + 1
	fmt.Printf("%d | %d | %d\n", min, max, winningSolutions)
}
