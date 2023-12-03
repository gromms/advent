package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
)

func processLine(symbolIndex int, line []byte, lineData [][]int) int {
    sum := 0

    if lineData != nil {
        for _, lineMatch := range lineData {
            if lineMatch[0] == lineMatch[1] {
                continue
            }

            num := getAdjacentNum(symbolIndex, line, lineMatch)

            if num >= 0 {
                // fmt.Printf("Previous line match: %d | num: %d\n", lineMatch, num)
            }

            sum += num
        }
    }

    return sum
}

func getAdjacentNum(symbolIndex int, line []byte, lineMatch []int) int {
    sum := 0

    // If contains number in g1 and symbol is adjacent to the number
    if lineMatch[3]-lineMatch[2] != 0 &&
        lineMatch[2]-1 <= symbolIndex && lineMatch[3] >= symbolIndex {
        num, _ := strconv.Atoi(string(line[lineMatch[2]:lineMatch[3]]))
        println(num)
        sum += num
    }

    // If contains number in g3 and symbol is adjacent to the number
    if lineMatch[6]-lineMatch[7] != 0 &&
        lineMatch[6]-1 <= symbolIndex && lineMatch[7] >= symbolIndex {
        num, _ := strconv.Atoi(string(line[lineMatch[6]:lineMatch[7]]))
        println(num)
        sum += num
    }

    return sum
}

func main() {

    file, _ := os.Open(os.Args[1])
    scanner := bufio.NewScanner(file)

    var previousLine []byte
    var previousLineMatches [][]int
    sum := 0

    inLineNumRe := regexp.MustCompile("([0-9]*)([^0-9.]*)([0-9]*)")

    for scanner.Scan() {
        println("Lines: ")
        var line = scanner.Text()
        fmt.Println(string(previousLine))
        fmt.Println(line)
        lineMatches := inLineNumRe.FindAllSubmatchIndex([]byte(line), -1)

        for _, match := range lineMatches {
            if match[1] == match[0] {
                continue
            }

            symbolIndex := -1

            // Get symbol index if there is one
            if match[4] != match[5] {
                symbolIndex = match[4]

                if symbolIndex >= 0 {
                    sum += processLine(symbolIndex, previousLine, previousLineMatches)
                }

                // If number exists
                if match[2] != match[3] {
                    num, _ := strconv.Atoi(string(line[match[2]:match[3]]))
                    sum += num
                }

                // If number right of symbol
                if match[6] != match[7] {
                    num, _ := strconv.Atoi(string(line[match[6]:match[7]]))
                    sum += num
                }
            }

            for _, prevLineMatch := range previousLineMatches {
                if prevLineMatch[4] == prevLineMatch[5] {
                    continue
                }

                symbolIndex = prevLineMatch[4]

                if symbolIndex >= 0 {
                    // fmt.Printf("%d\n", prevLineMatch)
                    sum += getAdjacentNum(symbolIndex, []byte(line), match)
                }
            }

            // fmt.Printf("%d\n", match)
        }

        previousLineMatches = lineMatches
        previousLine = []byte(line)
    }

    fmt.Printf("Result: %d", sum)
}
