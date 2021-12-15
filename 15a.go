package main

import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "fmt"
)

// assume row == column
func main() {
    file, _ := os.Open("15-sample.txt")
    defer file.Close()

    riskLevelMatrix := [][]int{}
    lowestRiskMatrix := [][]int{}

    scanner := bufio.NewScanner(file)
    size := 0
    for scanner.Scan() {
        riskLevelMatrix = append(riskLevelMatrix, []int{})
        lowestRiskMatrix = append(lowestRiskMatrix, []int{})

        for _, riskLevelChar := range strings.Split(scanner.Text(), "") {
            riskLevel, _ := strconv.Atoi(riskLevelChar)
            riskLevelMatrix[size] = append(riskLevelMatrix[size], riskLevel)
            lowestRiskMatrix[size] = append(lowestRiskMatrix[size], 0)
        }

        size += 1
    }
    // fmt.Println(riskLevelMatrix)

    lowestRiskMatrix[size - 1][size - 1] = riskLevelMatrix[size - 1][size - 1]
    for x := size - 2; x >= 0; x-- {
        lowestRiskMatrix[size - 1][x] = riskLevelMatrix[size - 1][x] + lowestRiskMatrix[size - 1][x + 1]
    }
    for y := size - 2; y >= 0; y-- {
        lowestRiskMatrix[y][size - 1] = riskLevelMatrix[y][size - 1] + lowestRiskMatrix[y + 1][size - 1]
    }
    for y := size - 2; y >= 0; y-- {
        for x := size - 2; x >= 0; x-- {
            if lowestRiskMatrix[y][x + 1] < lowestRiskMatrix[y + 1][x] {
                lowestRiskMatrix[y][x] = riskLevelMatrix[y][x] + lowestRiskMatrix[y][x + 1]
            } else {
                lowestRiskMatrix[y][x] = riskLevelMatrix[y][x] + lowestRiskMatrix[y + 1][x]
            }
        }
    }
    fmt.Println(lowestRiskMatrix)

    fmt.Println("Lowest total risk: ", lowestRiskMatrix[0][0] - riskLevelMatrix[0][0])
}
