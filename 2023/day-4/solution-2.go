package main

import (
    "bufio"
    "os"
    "slices"
    "strconv"
    "strings"
)

func calculateCardPoints(ticket int, line string) int {
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

    matches := 0

    for _, numStr := range strings.Split(line[separatorIndex+1:], " ") {
        if numStr == "" {
            continue
        }

        parsedInt, _ := strconv.Atoi(numStr)
        if slices.Contains(winningNumbers, parsedInt) {
            matches++
        }
    }

    return matches
}

func main() {
    fileName := os.Args[1]
    file, _ := os.Open(fileName)
    scanner := bufio.NewScanner(file)

    ticketCount := 0

    var ticketsWon = make(map[int]int)

    for scanner.Scan() {
        line := scanner.Text()
        matches := calculateCardPoints(ticketCount, line)

        ticketsWon[ticketCount] += 1

        for j := 1; j <= matches; j++ {
            wonCard := ticketCount + j
            ticketsWon[wonCard] += ticketsWon[ticketCount]
        }
        ticketCount++
    }

    sum := 0

    for i := 0; i <= ticketCount; i++ {
        sum += ticketsWon[i]
    }

    println(sum)
}
