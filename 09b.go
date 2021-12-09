package main

import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "sort"
    "fmt"
)

func recursivelySetBasinIndex(
    basinIndexMap map[int]map[int]int,
    currentBasinIndex int,
    x int,
    y int,
) {
    if basinIndexMap[y][x] != -1 { // height = 9 or already set
        return
    }

    // fmt.Println("set x:", x, ", y:", y, ", to ", currentBasinIndex)
    basinIndexMap[y][x] = currentBasinIndex
    if x > 0 {
        recursivelySetBasinIndex(basinIndexMap, currentBasinIndex, x - 1, y)
    }
    if y > 0 {
        recursivelySetBasinIndex(basinIndexMap, currentBasinIndex, x, y - 1)
    }
    if x < len(basinIndexMap[y]) - 1 {
        recursivelySetBasinIndex(basinIndexMap, currentBasinIndex, x + 1, y)
    }
    if y < len(basinIndexMap) - 1 {
        recursivelySetBasinIndex(basinIndexMap, currentBasinIndex, x, y + 1)
    }
}

func main() {
    file, _ := os.Open("09.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    basinIndexMap := make(map[int]map[int]int)
    line := 0
    for scanner.Scan() {
        basinIndexMap[line] = make(map[int]int)
        for x, heightString := range strings.Split(scanner.Text(), "") {
            height, _ := strconv.Atoi(heightString)
            if height == 9 {
                basinIndexMap[line][x] = -2
            } else {
                basinIndexMap[line][x] = -1
            }
        }
        line += 1
    }
    // fmt.Println(basinIndexMap)

    lowestBasinIndex := 0
    x := 0
    y := 0
    for {
        // fmt.Println("search x=", x, ", y=", y)

        recursivelySetBasinIndex(basinIndexMap, lowestBasinIndex, x, y)
        // fmt.Println(basinIndexMap)

        // find next
        xNew := x
        yNew := y
        for ; yNew < len(basinIndexMap); yNew++ {
            foundY := false
            for ; xNew < len(basinIndexMap[yNew]); xNew++ {
                if basinIndexMap[yNew][xNew] == -1 {
                    foundY = true
                    x = xNew
                    y = yNew
                    lowestBasinIndex += 1
                    break
                }
            }
            if foundY {
                break
            }
            xNew = 0
        }

        if yNew == len(basinIndexMap) && xNew == 0 {
            // all done
            break
        }
    }

    // fmt.Println(basinIndexMap)

    basinSizeMap := make(map[int]int)
    for _, basinIndexLine := range basinIndexMap {
        for _, basinIndex := range basinIndexLine {
            if basinIndex >= 0 {
                _, ok := basinSizeMap[basinIndex]
                if !ok {
                    basinSizeMap[basinIndex] = 0
                }
                basinSizeMap[basinIndex] += 1
            }
        }
    }
    // fmt.Println(basinSizeMap)

    basinSizes := make([]int, 0, len(basinSizeMap))
    for  _, value := range basinSizeMap {
        basinSizes = append(basinSizes, value)
    }

	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
    // fmt.Println(basinSizes)

    fmt.Println("Multiply three largest basin sizes:", basinSizes[0] * basinSizes[1] * basinSizes[2])
}
