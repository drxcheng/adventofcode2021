package main

import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "fmt"
)

func main() {
    file, _ := os.Open("09.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    heightMap := make(map[int]map[int]int)
    line := 0
    for scanner.Scan() {
        heightMap[line] = make(map[int]int)
        for x, heightString := range strings.Split(scanner.Text(), "") {
            height, _ := strconv.Atoi(heightString)
            heightMap[line][x] = height
        }
        line += 1
    }
    // fmt.Println(heightMap)

    riskLevels := 0
    for y, heightLine := range heightMap {
        for x, height := range heightLine {
            if (y == 0 || height < heightMap[y - 1][x]) &&
                (y == len(heightMap) - 1 || height < heightMap[y + 1][x]) &&
                (x == 0 || height < heightLine[x - 1]) &&
                (x == len(heightLine) - 1 || height < heightLine[x + 1]) {
                // found low point
                // fmt.Println("x=", x, ", y=", y, ", height=", height)
                riskLevels += height + 1
            }
        }
    }

    fmt.Println("Sum of the risk levels:", riskLevels)
}
