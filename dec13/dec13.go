package dec13

import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "fmt"
)

func Helper(onlyOneStep bool) {
    file, _ := os.Open("13.txt")
    defer file.Close()

    dotsMap := make(map[int]map[int]bool)
    scanner := bufio.NewScanner(file)
    maxX := 0
    maxY := 0
    for scanner.Scan() {
        line := scanner.Text()
        if line == "" {
            // end of coordinate
            break
        }
        coordinate := strings.Split(line, ",")
        x, _ := strconv.Atoi(coordinate[0])
        y, _ := strconv.Atoi(coordinate[1])
        _, ok := dotsMap[y]
        if !ok {
            dotsMap[y] = make(map[int]bool)
        }
        dotsMap[y][x] = true
        if x > maxX {
            maxX = x
        }
        if y > maxY {
            maxY = y
        }
    }
    // fmt.Println(dotsMap)

    for scanner.Scan() {
        instruction := scanner.Text()
        instructionParts := strings.Split(instruction, " ")
        foldInstruction := strings.Split(instructionParts[2], "=")
        along, _ := strconv.Atoi(foldInstruction[1])

        if foldInstruction[0] == "y" {
            for y := maxY; y > along; y-- {
                for x := 0; x <= maxX; x++ {
                    isDot, _ := dotsMap[y][x]
                    if isDot {
                        newY := along * 2 - y
                        _, ok := dotsMap[newY]
                        if !ok {
                            dotsMap[newY] = make(map[int]bool)
                        }
                        dotsMap[newY][x] = true
                    }
                }
            }
            // remove rows > along
            for y := maxY; y >= along; y-- {
                delete(dotsMap, y)
            }
            maxY = along - 1
        } else {
            for x := maxX; x > along; x-- {
                for y := 0; y <= maxY; y++ {
                    isDot, _ := dotsMap[y][x]
                    if isDot {
                        newX := along * 2 - x
                        dotsMap[y][newX] = true
                    }
                }
            }
            // remove columns > along
            for y, _ := range dotsMap {
                for x, _ := range dotsMap[y] {
                    if x >= along {
                        delete(dotsMap[y], x)
                    }
                }
            }
            maxX = along - 1
        }

        if onlyOneStep {
            break
        }
    }
    // fmt.Println(dotsMap)

    if onlyOneStep {
        // count
        totalDots := 0
        for y, _ := range dotsMap {
            totalDots += len(dotsMap[y])
        }
        fmt.Println("Total dots: ", totalDots)
    } else {
        // draw
        for y := 0; y <= maxY; y++ {
            line := ""
            for x := 0; x <= maxX; x++ {
                isDot, _ := dotsMap[y][x]
                if isDot {
                    line += "#"
                } else {
                    line += "."
                }
            }
            fmt.Println(line)
        }
    }
}
