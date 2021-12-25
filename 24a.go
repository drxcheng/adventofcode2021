package main

import (
    "os"
    "bufio"
    // "strconv"
    // "strings"
    "fmt"
)

func main() {
    file, _ := os.Open("24.txt")
    defer file.Close()

    // aSlice := []string{}
    // aMatrix := [][]string{}
    // aMap := make(map[int]map[int]string)

    scanner := bufio.NewScanner(file)
    // lineNumber := 0
    scanner.Scan()
    line := scanner.Text()
    for {
        newInstruction := false
        for scanner.Scan() {
            line = scanner.Text()
            fmt.Println(line)

            if string(line[0:3]) == "inp" {
                // fmt.Println(line)
                newInstruction = true
                break
            }
        }

        if !newInstruction {
            // end
            break
        }

        // for _, char := range strings.Split(line, "") {
        //     number, _ := strconv.Atoi(char)

        //     aSlice = append(aSlice, char)

        //     aMatrix[lineNumber] = append(aMatrix[lineNumber], char)

        //     _, ok := aMap[number]
        //     if !ok {
        //         aMap[number] = make(map[int]string)
        //     }
        //     aMap[number][0] = char

        // }

        // lineNumber += 1
    }
}
