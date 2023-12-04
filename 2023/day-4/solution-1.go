package main

import (
    "bufio"
    "os"
    "slices"
    "strconv"
    "strings"
)

func calculateCardPoints(line string, c chan int) {
    numStart := strings.Index(line, ":")
    separatorIndex := strings.Index(line, "|")

    var winningNumbers []int

    for _, numStr := range strings.Split(line[numStart+1:separatorIndex], " ") {
        if numStr == "" {
            continue
        }
        parsedInt, _ := strconv.Atoi(numStr)
        winningNumbers = append(winningNumbers, parsedInt)
    }

    // fmt.Printf("%d\n", winningNumbers)

    points := 0

    for _, numStr := range strings.Split(line[separatorIndex+1:], " ") {
        if numStr == "" {
            continue
        }

        parsedInt, _ := strconv.Atoi(numStr)
        if slices.Contains(winningNumbers, parsedInt) {
            if points == 0 {
                points = 1
            } else {
                points <<= 1
            }
        }
    }
    c <- points
}

func main() {
    fileName := os.Args[1]
    file, _ := os.Open(fileName)
    scanner := bufio.NewScanner(file)

    c := make(chan int)

    lines := 0

    for scanner.Scan() {
        line := scanner.Text()
        lines++
        go calculateCardPoints(line, c)
    }

    sum := 0

    for i := 0; i < lines; i++ {
        sum += <-c
    }

    println(sum)
}
