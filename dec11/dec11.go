package dec11

import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "fmt"
)

func recursiveLightUp(matrix [][]int, y int, x int) {
    lightUp := false
    if matrix[y][x] == 9 {
        lightUp = true
    }

    matrix[y][x] += 1
    if (!lightUp) {
        return
    }

    if y > 0 {
        recursiveLightUp(matrix, y - 1, x)
        if x > 0 {
            recursiveLightUp(matrix, y - 1, x - 1)
        }
        if x < len(matrix[y]) - 1 {
            recursiveLightUp(matrix, y - 1, x + 1)
        }
    }
    if x > 0 {
        recursiveLightUp(matrix, y, x - 1)
    }
    if x < len(matrix[y]) - 1 {
        recursiveLightUp(matrix, y, x + 1)
    }
    if y < len(matrix) - 1 {
        recursiveLightUp(matrix, y + 1, x)
        if x > 0 {
            recursiveLightUp(matrix, y + 1, x - 1)
        }
        if x < len(matrix[y]) - 1 {
            recursiveLightUp(matrix, y + 1, x + 1)
        }
    }
}

func Helper(maxStep int) {
    file, _ := os.Open("11.txt")
    defer file.Close()

    energyMatrix := [][]int{}
    scanner := bufio.NewScanner(file)
    line := 0
    for scanner.Scan() {
        energyMatrix = append(energyMatrix, []int{})
        for _, char := range strings.Split(scanner.Text(), "") {
            energy, _ := strconv.Atoi(char)
            energyMatrix[line] = append(energyMatrix[line], energy)
        }
        line += 1
    }
    // fmt.Println(energyMatrix)

    totalNumberOfFlashes := 0
    totalOctopuses := len(energyMatrix) * len(energyMatrix[0])
    for step := 0; ; step++ {
        for y := 0; y < len(energyMatrix); y++ {
            for x := 0; x < len(energyMatrix[y]); x++ {
                recursiveLightUp(energyMatrix, y, x)
            }
        }

        numberOfFlashes := 0
        for y := 0; y < len(energyMatrix); y++ {
            for x := 0; x < len(energyMatrix[y]); x++ {
                if energyMatrix[y][x] > 9 {
                    numberOfFlashes += 1
                    energyMatrix[y][x] = 0 // reset
                }
            }
        }

        if maxStep == 0 && numberOfFlashes == totalOctopuses {
            // part TWO
            fmt.Println("First step with simultaneously flash: ", step + 1)
            break
        }

        if maxStep > 0 && step >= maxStep {
            // part ONE
            fmt.Println("Total flahses: ", totalNumberOfFlashes)
            break
        }

        totalNumberOfFlashes += numberOfFlashes
    }
}
