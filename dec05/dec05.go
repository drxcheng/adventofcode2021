package dec05

import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "fmt"
    "math"
)

func Helper(allowDiagonal bool) {
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

        var xDelta int
        var yDelta int
        if x1 == x2 {
            // vertical
            xDelta = 0
            yDelta = (y2 - y1) / int(math.Abs(float64(y2 - y1)))
        } else if y1 == y2 {
            // horizontal
            xDelta = (x2 - x1) / int(math.Abs(float64(x2 - x1)))
            yDelta = 0
        } else if allowDiagonal && (x1 - x2 == y1 - y2 || x1 - x2 == y2 - y1) {
            // diagonal
            xDelta = (x2 - x1) / int(math.Abs(float64(x2 - x1)))
            yDelta = (y2 - y1) / int(math.Abs(float64(y2 - y1)))
        } else {
            continue
        }

        loop := 0
        for {
            x := x1 + xDelta * loop
            y := y1 + yDelta * loop

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

            if (xDelta > 0 && x >= x2) || (xDelta < 0 && x <= x2) ||
                (yDelta > 0 && y >= y2) || (yDelta < 0 && y <= y2) {
                break
            }

            loop += 1
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
