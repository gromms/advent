package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseNumbers(line string, sep string) []int {

	var nums []int
	a := strings.Split(line, sep)[1]
	for _, s := range strings.Split(a, " ") {
		if s != "" {
			num, _ := strconv.Atoi(s)
			nums = append(nums, num)
		}
	}
	return nums
}

func main() {

	fileName := os.Args[1]

	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)

	scanner.Scan()

	durations := parseNumbers(scanner.Text(), "Time:")
	fmt.Printf("%d\n", durations)

	scanner.Scan()

	distances := parseNumbers(scanner.Text(), "Distance:")
	fmt.Printf("%d\n", distances)

    result := 1

	for i := 0; i < len(durations); i++ {
		duration := durations[i]
		distance := distances[i] + 1

		halfSqrt := math.Sqrt(float64(duration*duration-4*distance)) / 2
		halfDuration := float64(duration) / 2

		min := int(math.Ceil(halfDuration - halfSqrt))
		max := int(math.Floor(halfDuration + halfSqrt))

        winningSolutions := max-min+1
		fmt.Printf("%d | %d | %d\n", min, max, winningSolutions)

        result *= winningSolutions
	}

    println(result)
}
