package main

import (
    "os"
    "bufio"
    "strings"
    "sort"
    "fmt"
)

func main() {
    file, _ := os.Open("10.txt")
    defer file.Close()

    CHUNK_MAP := map[string]string{
        "(": ")",
        "[": "]",
        "{": "}",
        "<": ">",
    }
    SCORE := map[string]int {
        "(": 1,
        "[": 2,
        "{": 3,
        "<": 4,
    }

    scanner := bufio.NewScanner(file)
    autocompleteScores := []int{}
    for scanner.Scan() {
        chunkStack := []string{}
        corrupted := false
        for _, char := range strings.Split(scanner.Text(), "") {
            _, ok := CHUNK_MAP[char]
            if ok {
                // left
                chunkStack = append(chunkStack, char)
            } else {
                // right, get stack value
                lastLeft := chunkStack[len(chunkStack) - 1]
                expectedRight, _ := CHUNK_MAP[lastLeft]
                if expectedRight != char {
                    // corrupted, ignore this line
                    corrupted = true
                    break
                }
                chunkStack = chunkStack[:len(chunkStack) - 1]
            }
        }

        if corrupted {
            continue
        }

        autocompleteScore := 0
        for len(chunkStack) > 0 {
            lastLeft := chunkStack[len(chunkStack) - 1]
            chunkStack = chunkStack[:len(chunkStack) - 1]
            charScore, _ := SCORE[lastLeft]
            autocompleteScore = autocompleteScore * 5 + charScore
        }
        autocompleteScores = append(autocompleteScores, autocompleteScore)
    }
    sort.Ints(autocompleteScores)
    // fmt.Println(autocompleteScores)

    fmt.Println("Middle score:", autocompleteScores[(len(autocompleteScores) - 1)/2])
}
