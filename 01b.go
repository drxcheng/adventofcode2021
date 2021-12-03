package main

import (
    "os"
    "log"
    "bufio"
    "strconv"
    "fmt"
)

func main() {
    file, err := os.Open("01.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var depths []int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        depth, err := strconv.Atoi(scanner.Text())
        if err != nil {
            log.Fatal(err)
        }
        depths = append(depths, depth)
    }

    var numberOfIncreases int = 0
    var prevSlidingWindowValue int = depths[0] + depths[1] + depths[2]
    var index int = 3
    for index < len(depths) {
        var slidingWindowValue int = depths[index - 2] + depths[index - 1] + depths[index]
        if (slidingWindowValue > prevSlidingWindowValue) {
            numberOfIncreases += 1
        }
        prevSlidingWindowValue = slidingWindowValue
        index += 1
    }

    fmt.Println("number of increases: ", numberOfIncreases)
}
