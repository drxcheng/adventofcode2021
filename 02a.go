package main

import (
    "os"
    "log"
    "bufio"
    "strconv"
    "strings"
    "math"
    "fmt"
)

func main() {
    file, err := os.Open("02.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var offsetHorizontal int = 0
    var offsetVertical int = 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lineData := strings.Split(scanner.Text(), " ")
        direction := lineData[0]
        amount, err := strconv.Atoi(lineData[1])
        if err != nil {
            log.Fatal(err)
        }

        if (direction == "forward") {
            offsetHorizontal += amount
        } else if (direction == "down") {
            offsetVertical += amount
        } else if (direction == "up") {
            offsetVertical -= amount
        } else {
            log.Fatal("Wrong direction in line ", lineData)
        }
    }

    fmt.Println("horizontal position: ", offsetHorizontal)
    fmt.Println("vertical position: ", offsetVertical)
    fmt.Println("result: ", int(math.Abs(float64(offsetHorizontal * offsetVertical))))
}
