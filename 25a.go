package main

import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "fmt"
)

func main() {
    file, _ := os.Open("25.txt")
    defer file.Close()

    seaMatrix := [][]string{}

    scanner := bufio.NewScanner(file)
    lineNumber := 0
    for scanner.Scan() {
        seaMatrix = append(seaMatrix, []string{})
        line := scanner.Text()
        for _, char := range strings.Split(line, "") {
            if char == "v" || char == ">" {
                char = char + strconv.Itoa(0)
            }
            seaMatrix[lineNumber] = append(seaMatrix[lineNumber], char)
        }
        lineNumber += 1
    }
    fmt.Println(seaMatrix)

    step := 1
    for {
        fmt.Println("Step", step)
        hasMove := false
        for y := 0; y < len(seaMatrix); y++ {
            originalLeftMost := seaMatrix[y][0]
            for x := 0; x < len(seaMatrix[y]); x++ {
                if seaMatrix[y][x] == ">" + strconv.Itoa(step - 1) {
                    if x < len(seaMatrix[y]) - 1 && seaMatrix[y][x + 1] == "." {
                        seaMatrix[y][x] = "."
                        seaMatrix[y][x + 1] = ">" + strconv.Itoa(step)
                        hasMove = true
                    } else if x == len(seaMatrix[y]) - 1 && originalLeftMost == "." {
                        seaMatrix[y][x] = "."
                        seaMatrix[y][0] = ">" + strconv.Itoa(step)
                        hasMove = true
                    } else {
                        seaMatrix[y][x] = ">" + strconv.Itoa(step)
                    }
                }
            }
        }
        for x := 0; x < len(seaMatrix[0]); x++ {
            originalUpMost := seaMatrix[0][x]
            for y := 0; y < len(seaMatrix); y++ {
                if seaMatrix[y][x] == "v" + strconv.Itoa(step - 1) {
                    if y < len(seaMatrix) - 1 && seaMatrix[y + 1][x] == "." {
                        seaMatrix[y][x] = "."
                        seaMatrix[y + 1][x] = "v" + strconv.Itoa(step)
                        hasMove = true
                    } else if y == len(seaMatrix) - 1 && originalUpMost == "." {
                        seaMatrix[y][x] = "."
                        seaMatrix[0][x] = "v" + strconv.Itoa(step)
                        hasMove = true
                    } else {
                        seaMatrix[y][x] = "v" + strconv.Itoa(step)
                    }
                }
            }
        }

        if !hasMove {
            fmt.Println("Stop at step", step)
            break
        }

        step += 1
    }

    // fmt.Println(seaMatrix)
}
