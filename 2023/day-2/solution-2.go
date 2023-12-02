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

    numStart := -1

    sum := 0

    highestColorsInGame := map[rune]int{
        'r': 0,
        'g': 0,
        'b': 0,
    }

    state := game

    for index, element := range file {
        char := rune(element)

        switch state {
        case game:
            if char == ':' {
                state = color
                numStart = -1
                highestColorsInGame['r'] = 0
                highestColorsInGame['g'] = 0
                highestColorsInGame['b'] = 0
            }
            break
        case color:
            if (char == 'r' || char == 'g' || char == 'b') && numStart > 0 {
                num, _ := getInt(file, numStart, index-1)

                highestColorsInGame[char] = max(highestColorsInGame[char], num)
                numStart = -1
            } else if isNum(element) && numStart < 0 {
                numStart = index
            } else if char == '\n' {
                state = game
                sum += highestColorsInGame['r'] * highestColorsInGame['g'] * highestColorsInGame['b']
            }
            break
        }
    }

    fmt.Printf("result: %d", sum)

}
