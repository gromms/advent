package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
)

func processLine(symbolLine int, symbolIndex int, line []byte, lineData [][]int, gears map[string][]int) {
    if lineData != nil {
        for _, lineMatch := range lineData {
            if lineMatch[0] == lineMatch[1] {
                continue
            }

            getAdjacentNum(symbolLine, symbolIndex, line, lineMatch, gears)
        }
    }

}

func getGearKey(line int, index int) string {
    return strconv.Itoa(line) + ":" + strconv.Itoa(index)
}

func getAdjacentNum(symbolLine int, symbolIndex int, line []byte, lineMatch []int, gears map[string][]int) {
    // If contains number in g1 and symbol is adjacent to the number
    if lineMatch[3]-lineMatch[2] != 0 &&
        lineMatch[2]-1 <= symbolIndex && lineMatch[3] >= symbolIndex {
        num, _ := strconv.Atoi(string(line[lineMatch[2]:lineMatch[3]]))
        key := getGearKey(symbolLine, symbolIndex)
        gear, ok := gears[key]
        if !ok {
            var gear []int
            gears[key] = gear
        }

        gears[key] = append(gear, num)
    }

    // If contains number in g3 and symbol is adjacent to the number
    if lineMatch[6]-lineMatch[7] != 0 &&
        lineMatch[6]-1 <= symbolIndex && lineMatch[7] >= symbolIndex {
        num, _ := strconv.Atoi(string(line[lineMatch[6]:lineMatch[7]]))
        key := getGearKey(symbolLine, symbolIndex)
        gear, ok := gears[key]
        if !ok {
            var gear []int
            gears[key] = gear
        }

        gears[key] = append(gear, num)
    }
}

func main() {

    file, _ := os.Open(os.Args[1])
    scanner := bufio.NewScanner(file)

    var previousLine []byte
    var previousLineMatches [][]int
    sum := 0

    inLineNumRe := regexp.MustCompile("([0-9]*)([*]*)([0-9]*)")

    gears := map[string][]int{}
    lineNum := 0

    for ; scanner.Scan(); lineNum++ {
        var line = scanner.Text()
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
                    processLine(lineNum, symbolIndex, previousLine, previousLineMatches, gears)
                }

                // If number exists
                if match[2] != match[3] {
                    num, _ := strconv.Atoi(line[match[2]:match[3]])
                    key := getGearKey(lineNum, symbolIndex)
                    gear, ok := gears[key]
                    if !ok {
                        var gear []int
                        gears[key] = gear
                    }

                    gears[key] = append(gear, num)
                }

                // If number right of symbol
                if match[6] != match[7] {
                    num, _ := strconv.Atoi(line[match[6]:match[7]])
                    key := getGearKey(lineNum, symbolIndex)
                    gear, ok := gears[key]
                    if !ok {
                        var gear []int
                        gears[key] = gear
                    }

                    gears[key] = append(gear, num)
                }
            }

            for _, prevLineMatch := range previousLineMatches {
                if prevLineMatch[4] == prevLineMatch[5] {
                    continue
                }

                symbolIndex = prevLineMatch[4]

                if symbolIndex >= 0 {
                    getAdjacentNum(lineNum-1, symbolIndex, []byte(line), match, gears)
                }
            }
        }

        previousLineMatches = lineMatches
        previousLine = []byte(line)
    }

    println(gears)

    for key, val := range gears {
        fmt.Printf("%s : %d\n", key, val)
        if len(val) != 2 {
            continue
        }

        sum += val[0] * val[1]
    }

    fmt.Printf("Result: %d", sum)
}
