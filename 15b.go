package main

import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "math"
    "fmt"
)

const REPEAT = 5

// assume row == column
// not working
func getRiskLevel(riskLevelMatrix [][]int, size int, x int, y int) int {
    xTimes := int(math.Floor(float64(x / size)))
    xMod := x % size
    yTimes := int(math.Floor(float64(y / size)))
    yMod := y % size

    riskLevel := riskLevelMatrix[yMod][xMod]
    riskLevel += xTimes + yTimes
    if riskLevel > 9 {
        return riskLevel % 9
    }
    return riskLevel
}

func main() {
    file, _ := os.Open("15.txt")
    defer file.Close()

    riskLevelMatrix := [][]int{}

    scanner := bufio.NewScanner(file)
    size := 0
    for scanner.Scan() {
        riskLevelMatrix = append(riskLevelMatrix, []int{})

        for _, riskLevelChar := range strings.Split(scanner.Text(), "") {
            riskLevel, _ := strconv.Atoi(riskLevelChar)
            riskLevelMatrix[size] = append(riskLevelMatrix[size], riskLevel)
        }

        size += 1
    }
    // fmt.Println(riskLevelMatrix)

    lowestRiskMap := make(map[int]map[int]int)

    _, ok := lowestRiskMap[size * REPEAT - 1]
    if !ok {
        lowestRiskMap[size * REPEAT - 1] = make(map[int]int)
    }
    lowestRiskMap[size * REPEAT - 1][size * REPEAT - 1] = getRiskLevel(riskLevelMatrix, size, size * REPEAT - 1, size * REPEAT - 1)
    for y := size * REPEAT - 2; y >= 0; y-- {
        lowestRiskMap[y] = make(map[int]int)
        lowestRiskMap[y][size * REPEAT - 1] = getRiskLevel(riskLevelMatrix, size, size * REPEAT - 1, y) + lowestRiskMap[y + 1][size * REPEAT - 1]
    }
    for x := size * REPEAT - 2; x >= 0; x-- {
        lowestRiskMap[size * REPEAT - 1][x] = getRiskLevel(riskLevelMatrix, size, x, size * REPEAT - 1) + lowestRiskMap[size * REPEAT - 1][x + 1]
    }

    for y := size * REPEAT - 2; y >= 0; y-- {
        for x := size * REPEAT - 2; x >= 0; x-- {
            if lowestRiskMap[y][x + 1] < lowestRiskMap[y + 1][x] {
                lowestRiskMap[y][x] = getRiskLevel(riskLevelMatrix, size, x, y) + lowestRiskMap[y][x + 1]
            } else {
                lowestRiskMap[y][x] = getRiskLevel(riskLevelMatrix, size, x, y) + lowestRiskMap[y + 1][x]
            }
        }
    }
    fmt.Println(lowestRiskMap)

    fmt.Println("Lowest total risk: ", lowestRiskMap[0][0] - getRiskLevel(riskLevelMatrix, size, 0, 0))
}
