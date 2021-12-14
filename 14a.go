package main

import (
    "os"
    "bufio"
    "strings"
    "sort"
    "fmt"
)

const STEP = 40

func main() {
    file, _ := os.Open("14-sample.txt")
    defer file.Close()

    insertionRulesMap := make(map[string]string)

    scanner := bufio.NewScanner(file)
    scanner.Scan()
    polymer := scanner.Text()
    scanner.Scan()

    for scanner.Scan() {
        insertionLineParts := strings.Split(scanner.Text(), " -> ")

        insertionRulesMap[insertionLineParts[0]] = insertionLineParts[1]
    }

    // fmt.Println("Polymer template:", polymer)
    // fmt.Println("insertionRulesMap:", insertionRulesMap)

    for step := 0; step < STEP; step++ {
        fmt.Println("step", step)
        for i := len(polymer) - 2; i >= 0; i-- {
            pair := polymer[i:i+2]
            insertion, found := insertionRulesMap[pair]
            if found {
                // fmt.Println("insert", insertion, " at index", i+1)
                polymer = polymer[:i+1] + insertion + polymer[i+1:]
            }
        }

        // fmt.Println("After step", step + 1, ":", polymer)
        // fmt.Println("After step", step + 1, ", length:", len(polymer))
    }

    // count
    charCountMap := make(map[string]int)
    for _, char := range strings.Split(polymer, "") {
        _, found := charCountMap[char]
        if !found {
            charCountMap[char] = 0
        }

        charCountMap[char] += 1
    }
    fmt.Println(charCountMap)

    sortedCount := []int{}
    for _, count := range charCountMap {
        sortedCount = append(sortedCount, count)
    }
    sort.Ints(sortedCount)
    fmt.Println(sortedCount)

    fmt.Println("result =", sortedCount[len(sortedCount) - 1] - sortedCount[0])
}
