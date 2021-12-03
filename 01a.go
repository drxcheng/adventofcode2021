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

    var numberOfIncreases int = 0
    var prevDepth int = 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        depth, err := strconv.Atoi(scanner.Text())
        if err != nil {
            log.Fatal(err)
        }
        if (prevDepth > 0 && depth > prevDepth) {
            numberOfIncreases += 1
        }
        prevDepth = depth
    }

    fmt.Println("number of increases: ", numberOfIncreases)
}
