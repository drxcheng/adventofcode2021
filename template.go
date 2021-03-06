package main

import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "fmt"
)

func main() {
    file, _ := os.Open("sample.txt")
    defer file.Close()

    aSlice := []string{}
    aMatrix := [][]string{}
    aMap := make(map[int]map[int]string)

    scanner := bufio.NewScanner(file)
    lineNumber := 0
    for scanner.Scan() {
        aMatrix = append(aMatrix, []string{})

        line := scanner.Text()

        for _, char := range strings.Split(line, "") {
            number, _ := strconv.Atoi(char)

            aSlice = append(aSlice, char)

            aMatrix[lineNumber] = append(aMatrix[lineNumber], char)

            _, ok := aMap[number]
            if !ok {
                aMap[number] = make(map[int]string)
            }
            aMap[number][0] = char

        }

        lineNumber += 1
    }

    fmt.Println("slice:", aSlice)
    fmt.Println("matrix:", aMatrix)
    fmt.Println("map:", aMap)
}
