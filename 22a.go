package main

import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "fmt"
)

func setCubeStatus(cube map[int]map[int]map[int]bool, step string) map[int]map[int]map[int]bool {
    on := false
    if string(step[0:2]) == "on" {
        on = true
    }

    startIndex := strings.Index(step, "x=")
    endIndex := strings.Index(step, ",")
    valuesString := strings.Split(string(step[startIndex + 2 : endIndex]), "..")
    xStartValue, _ := strconv.Atoi(valuesString[0])
    xEndValue, _ := strconv.Atoi(valuesString[1])

    step = step[endIndex + 1:]
    startIndex = strings.Index(step, "y=")
    endIndex = strings.Index(step, ",")
    valuesString = strings.Split(string(step[startIndex + 2 : endIndex]), "..")
    yStartValue, _ := strconv.Atoi(valuesString[0])
    yEndValue, _ := strconv.Atoi(valuesString[1])

    step = step[endIndex + 1:]
    startIndex = strings.Index(step, "z=")
    endIndex = len(step)
    valuesString = strings.Split(string(step[startIndex + 2 : endIndex]), "..")
    zStartValue, _ := strconv.Atoi(valuesString[0])
    zEndValue, _ := strconv.Atoi(valuesString[1])

    // fmt.Println(xStartValue, "-", xEndValue, ",", yStartValue, "-", yEndValue, ",", zStartValue, "-", zEndValue)
    for x := xStartValue; x <= xEndValue; x++ {
        if x < -50 || x > 50 {
            continue
        }
        for y := yStartValue; y <= yEndValue; y++ {
            if y < -50 || y > 50 {
                continue
            }
            for z := zStartValue; z <= zEndValue; z++ {
                if z < -50 || z > 50 {
                    continue
                }
                _, xok := cube[x]
                if !xok {
                    cube[x] = make(map[int]map[int]bool)
                }
                _, yok := cube[x][y]
                if !yok {
                    cube[x][y] = make(map[int]bool)
                }

                cube[x][y][z] = on
            }
        }
    }

    return cube
}

func countCube(cube map[int]map[int]map[int]bool) int {
    numberOfOn := 0
    for xIndex, _ := range cube {
        for yIndex, _ := range cube[xIndex] {
            for _, on := range cube[xIndex][yIndex] {
                if on {
                    numberOfOn += 1
                }
            }
        }
    }

    return numberOfOn
}

func main() {
    file, _ := os.Open("22.txt")
    defer file.Close()

    cubeMap := make(map[int]map[int]map[int]bool)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        cubeMap = setCubeStatus(cubeMap, scanner.Text())
    }
    // fmt.Println(cubeMap)

    numberOfOn := countCube(cubeMap)
    fmt.Println("Number of on:", numberOfOn)
}
