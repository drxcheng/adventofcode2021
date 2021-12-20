package main

import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "fmt"
)

func expandImage(imageMap map[int]map[int]rune) map[int]map[int]rune {
    expandedImageMap := make(map[int]map[int]rune)

    for i := 0; i < len(imageMap); i++ {
        _, ok := expandedImageMap[i + 1]
        if !ok {
            expandedImageMap[i + 1] = make(map[int]rune)
        }

        expandedImageMap[i + 1][0] = '0'
        for j := 0; j < len(imageMap[i]); j++ {
            expandedImageMap[i + 1][j + 1] = imageMap[i][j]
        }
        expandedImageMap[i + 1][len(imageMap[i]) + 1] = '0'
    }

    expandedImageMap[0] = make(map[int]rune)
    expandedImageMap[len(imageMap) + 1] = make(map[int]rune)
    for i := 0; i < len(imageMap[0]) + 2; i++ {
        expandedImageMap[0][i] = '0'
        expandedImageMap[len(expandedImageMap) - 1][i] = '0'
    }

    return expandedImageMap
}

func enhanceImage(algorithm []string, imageInputMap map[int]map[int]rune) map[int]map[int]rune {
    enhancedImageMap := make(map[int]map[int]rune)

    for row := 0; row < len(imageInputMap); row++ {
        _, ok := enhancedImageMap[row]
        if !ok {
            enhancedImageMap[row] = make(map[int]rune)
        }

        for column := 0; column < len(imageInputMap[row]); column++ {
            var binaryString string
            if row == 0 && column == 0 {
                binaryString = "0000" +
                    string(imageInputMap[row][column]) +
                    string(imageInputMap[row][column + 1]) +
                    "0" +
                    string(imageInputMap[row + 1][column]) +
                    string(imageInputMap[row + 1][column + 1])
            } else if row == 0 && column == len(imageInputMap[row]) - 1 {
                binaryString = "000" +
                    string(imageInputMap[row][column - 1]) +
                    string(imageInputMap[row][column]) +
                    "0" +
                    string(imageInputMap[row + 1][column - 1]) +
                    string(imageInputMap[row + 1][column]) + "0"
            } else if row == len(imageInputMap) - 1 && column == 0 {
                binaryString = "0" +
                    string(imageInputMap[row - 1][column]) +
                    string(imageInputMap[row - 1][column + 1]) +
                    "0" +
                    string(imageInputMap[row][column]) +
                    string(imageInputMap[row][column + 1]) + "000"
            } else if row == len(imageInputMap) - 1 && column == len(imageInputMap[row]) - 1 {
                binaryString =
                    string(imageInputMap[row - 1][column - 1]) +
                    string(imageInputMap[row - 1][column]) +
                    "0" +
                    string(imageInputMap[row][column - 1]) +
                    string(imageInputMap[row][column]) + "0000"
            } else if row == 0 {
                binaryString = "000" +
                    string(imageInputMap[row][column - 1]) +
                    string(imageInputMap[row][column]) +
                    string(imageInputMap[row][column + 1]) +
                    string(imageInputMap[row + 1][column - 1]) +
                    string(imageInputMap[row + 1][column]) +
                    string(imageInputMap[row + 1][column + 1])
            } else if row == len(imageInputMap) - 1 {
                binaryString =
                    string(imageInputMap[row - 1][column - 1]) +
                    string(imageInputMap[row - 1][column]) +
                    string(imageInputMap[row - 1][column + 1]) +
                    string(imageInputMap[row][column - 1]) +
                    string(imageInputMap[row][column]) +
                    string(imageInputMap[row][column + 1]) + "000"
            } else if column == 0 {
                binaryString = "0" +
                    string(imageInputMap[row - 1][column]) +
                    string(imageInputMap[row - 1][column + 1]) +
                    "0" +
                    string(imageInputMap[row][column]) +
                    string(imageInputMap[row][column + 1]) +
                    "0" +
                    string(imageInputMap[row + 1][column]) +
                    string(imageInputMap[row + 1][column + 1])
            } else if column == len(imageInputMap[row]) - 1 {
                binaryString =
                    string(imageInputMap[row - 1][column - 1]) +
                    string(imageInputMap[row - 1][column]) +
                    "0" +
                    string(imageInputMap[row][column - 1]) +
                    string(imageInputMap[row][column]) +
                    "0" +
                    string(imageInputMap[row + 1][column - 1]) +
                    string(imageInputMap[row + 1][column]) + "0"
            } else {
                binaryString =
                    string(imageInputMap[row - 1][column - 1]) +
                    string(imageInputMap[row - 1][column]) +
                    string(imageInputMap[row - 1][column + 1]) +
                    string(imageInputMap[row][column - 1]) +
                    string(imageInputMap[row][column]) +
                    string(imageInputMap[row][column + 1]) +
                    string(imageInputMap[row + 1][column - 1]) +
                    string(imageInputMap[row + 1][column]) +
                    string(imageInputMap[row + 1][column + 1])
            }

            binaryStringValue, _ := strconv.ParseInt(binaryString, 2, 64)
            if algorithm[binaryStringValue] == "#" {
                enhancedImageMap[row][column] = '1'
            } else {
                enhancedImageMap[row][column] = '0'
            }
        }
    }

    return enhancedImageMap
}

func printImage(imageMap map[int]map[int]rune) {
    for row := 0; row < len(imageMap); row++ {
        for column := 0; column < len(imageMap[row]); column++ {
            if imageMap[row][column] == '1' {
                fmt.Print("#")
            } else {
                fmt.Print(".")
            }
        }
        fmt.Println()
    }
}

func countLight(imageMap map[int]map[int]rune) int {
    light := 0
    for row := 0; row < len(imageMap); row++ {
        for column := 0; column < len(imageMap[row]); column++ {
            if imageMap[row][column] == '1' {
                light += 1
            }
        }
    }

    return light
}

func main() {
    file, _ := os.Open("20.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Scan()
    algorithm := strings.Split(scanner.Text(), "")
    // fmt.Println(algorithm)
    scanner.Scan()

    imageInputMap := make(map[int]map[int]rune)

    lineNumber := 0
    for scanner.Scan() {
        _, ok := imageInputMap[lineNumber]
        if !ok {
            imageInputMap[lineNumber] = make(map[int]rune)
        }

        for i, char := range strings.Split(scanner.Text(), "") {
            if char == "#" {
                imageInputMap[lineNumber][i] = '1'
            } else {
                imageInputMap[lineNumber][i] = '0'
            }
        }

        lineNumber += 1
    }
    // fmt.Println("input: ")
    // printImage(imageInputMap)

    var totalLights int
    var prevTotalLights int
    for {
        enhancedImageMap := enhanceImage(algorithm, imageInputMap)
        // fmt.Println("enhanced: ")
        // printImage(enhancedImageMap)

        enhancedImageMap = enhanceImage(algorithm, enhancedImageMap)
        // fmt.Println("further enhanced: ")
        // printImage(enhancedImageMap)

        totalLights = countLight(enhancedImageMap)
        fmt.Println("Total lights: ", totalLights, ", Previous total lights: ", prevTotalLights)

        if totalLights != prevTotalLights {
            prevTotalLights = totalLights
            imageInputMap = expandImage(imageInputMap)
            // fmt.Println("expanded input: ")
            // printImage(imageInputMap)
        } else {
            break
        }
    }

    fmt.Println("Total lights: ", totalLights)
}
