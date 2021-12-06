package main

import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "fmt"
)

func processOneDay(states []int) []int {
    numberOfNewFishes := 0
    for i, s := range states {
        if s == 0 {
            states[i] = 6
            numberOfNewFishes += 1
        } else {
            states[i] -= 1
        }
    }

    i := 0
    for numberOfNewFishes - i > 0 {
        states = append(states, 8)
        i += 1
    }

    return states
}

func main() {
    file, _ := os.Open("06-sample.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    fishStates := []int{}
    scanner.Scan()
    for _, v := range strings.Split(scanner.Text(), ",") {
        state, _ := strconv.Atoi(v)
        fishStates = append(fishStates, state)
    }

    // fmt.Println(fishStates)

    for i := 0; i < 28; i++ {
        fishStates = processOneDay(fishStates)
        fmt.Println(fishStates)
        fmt.Println("after day ", (i + 1), ", count=", len(fishStates))
        // if (i + 1) % 7 == 0 {
        // }
        // fmt.Println(fishStates)
    }

    fmt.Println("Number of fishes: ", len(fishStates))
}
