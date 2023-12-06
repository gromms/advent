package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

type Map struct {
    dest   int64
    source int64
    length int64
}

func main() {
    fileName := os.Args[1]

    file, _ := os.Open(fileName)
    scanner := bufio.NewScanner(file)

    var mappingBuf []int64

    scanner.Scan()
    line := scanner.Text()
    var mappingVals []Map

    for _, seed := range strings.Split(strings.Split(line, ":")[1], " ") {
        if seed == "" {
            continue
        }

        seedNum, _ := strconv.ParseInt(seed, 10, 64)
        mappingBuf = append(mappingBuf, seedNum)
    }

    fmt.Printf("%d\n", mappingBuf)

    scanner.Scan()
    for stage := 0; stage < 7; stage++ {
        scanner.Scan()

        mappingVals = make([]Map, 0)
        for scanner.Scan() {
            line = scanner.Text()
            if line == "" {
                break
            }

            vals := strings.Split(line, " ")
            dest, _ := strconv.ParseInt(vals[0], 10, 64)
            src, _ := strconv.ParseInt(vals[1], 10, 64)
            length, _ := strconv.ParseInt(vals[2], 10, 64)
            mappingVals = append(mappingVals, Map{dest, src, length})
        }
        // for _, mappingVal := range mappingVals {
        //     fmt.Printf("%d | %d | %d\n", mappingVal.dest, mappingVal.source, mappingVal.length)
        // }

        for i, bufVal := range mappingBuf {
            for _, mappingVal := range mappingVals {
                if bufVal >= mappingVal.source && bufVal <= mappingVal.source+mappingVal.length-1 {
                    offset := bufVal - mappingVal.source
                    mappingBuf[i] = mappingVal.dest + offset
                }
            }
        }
        fmt.Printf("%d\n", mappingBuf)

    }

    sort.Slice(mappingBuf, func(i, j int) bool {
        return mappingBuf[i] < mappingBuf[j]
    })

    fmt.Println(mappingBuf[0])
    fmt.Println(mappingBuf[len(mappingBuf)-1])

}
