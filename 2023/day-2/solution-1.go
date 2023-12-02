package main

import (
    "fmt"
    "io/fs"
    "os"
    "strconv"
)

type State uint8

const (
    game State = iota
    color
    skip
)

func getInt(text []byte, start int, end int) (int, error) {
    return strconv.Atoi(string(text[start:end]))
}

func isNum(b byte) bool {
    return b >= 48 && b <= 57
}

func main() {

    dir := os.DirFS("./")
    file, error := fs.ReadFile(dir, "input")

    if error != nil {
        panic(error)
    }

    var gameNumber int
    numStart := -1

    sum := 0

    limits := map[rune]int{
        'r': 12,
        'g': 13,
        'b': 14,
    }

    state := game

    for index, element := range file {
        char := rune(element)

        switch state {
        case game:
            if char == ':' {
                state = color
                gameNumber, _ = strconv.Atoi(string(file[numStart:index]))
                numStart = -1
            } else if numStart < 0 && isNum(element) {
                numStart = index
            }
            break
        case color:
            if (char == 'r' || char == 'g' || char == 'b') && numStart > 0 {
                num, _ := getInt(file, numStart, index-1)
                numStart = -1

                if num > limits[char] {
                    state = skip
                    break
                }
            } else if isNum(element) && numStart < 0 {
                numStart = index
            } else if char == '\n' {
                state = game
                // fmt.Printf("Game %d is possible\n", gameNumber)
                sum += gameNumber
            }
            break
        case skip:
            if char == '\n' {
                state = game
            }
            break
        }
    }

    fmt.Printf("result: %d", sum)

}
