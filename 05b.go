package main

import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "fmt"
)

func main() {
    file, _ := os.Open("05.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    f := func(c rune) bool {
        return c == ',' || c == ' ' || c == '-' || c == '>'
    }
    fieldMap := make(map[int]map[int]int)
    for scanner.Scan() {
        positions := strings.FieldsFunc(scanner.Text(), f)
        if len(positions) != 4 {
            continue
        }
        // fmt.Println(positions)
        x1, _ := strconv.Atoi(positions[0])
        y1, _ := strconv.Atoi(positions[1])
        x2, _ := strconv.Atoi(positions[2])
        y2, _ := strconv.Atoi(positions[3])

        if x1 != x2 && y1 != y2 && (x1 - x2) != (y1 - y2) && (x1 - x2) != (y2 - y1) {
            continue
        }

        if x1 == x2 {
            // vertical
            var yL int
            var yH int
            if y1 > y2 {
                yL = y2
                yH = y1
            } else {
                yL = y1
                yH = y2
            }

            _, ok := fieldMap[x1]
            if !ok {
                fieldMap[x1] = make(map[int]int)
            }
            for y := yL; y <= yH; y++ {
                _, ok := fieldMap[x1][y]
                if !ok {
                    fieldMap[x1][y] = 1
                } else {
                    fieldMap[x1][y] += 1
                }
            }
        } else if y1 == y2 {
            // horizontal
            var xL int
            var xH int
            if x1 > x2 {
                xL = x2
                xH = x1
            } else {
                xL = x1
                xH = x2
            }

            for x := xL; x <= xH; x++ {
                _, ok := fieldMap[x]
                if !ok {
                    fieldMap[x] = make(map[int]int)
                }
                _, ok = fieldMap[x][y1]
                if !ok {
                    fieldMap[x][y1] = 1
                } else {
                    fieldMap[x][y1] += 1
                }
            }
        } else {
            // diagnal
            var xL int
            var xH int
            var y int
            var yDelta int
            if x1 > x2 {
                xL = x2
                xH = x1
                y = y2
                if y1 > y2 {
                    yDelta = 1
                } else {
                    yDelta = -1
                }
            } else {
                xL = x1
                xH = x2
                y = y1
                if y1 > y2 {
                    yDelta = -1
                } else {
                    yDelta = 1
                }
            }

            for x := xL; x <= xH; x++ {
                y := y + (x - xL) * yDelta
                _, ok := fieldMap[x]
                if !ok {
                    fieldMap[x] = make(map[int]int)
                }
                _, ok = fieldMap[x][y]
                if !ok {
                    fieldMap[x][y] = 1
                } else {
                    fieldMap[x][y] += 1
                }
            }
        }
    }

    // fmt.Println(fieldMap)
    numberOfPoints := 0

    for x := range fieldMap {
        for _, v := range fieldMap[x] {
            if v > 1 {
                numberOfPoints += 1
            }
        }
    }

    fmt.Println("Number of points with at least two lines overlap: ", numberOfPoints)
}
