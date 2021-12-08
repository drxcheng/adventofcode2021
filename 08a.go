package main

import (
    "os"
    "bufio"
    "strings"
    "fmt"
)

func main() {
    file, _ := os.Open("08.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    numberOfSimpleDigits := 0
    for scanner.Scan() {
        data := strings.Split(scanner.Text(), " | ")
        for _, digit := range strings.Fields(data[1]) {
            numberOfSegments := len(digit)
            if numberOfSegments == 2 || numberOfSegments == 4 ||
                numberOfSegments == 3 || numberOfSegments == 7 {
                numberOfSimpleDigits += 1
            }
        }
    }

    fmt.Println("Number of simple digits:", numberOfSimpleDigits)
}
