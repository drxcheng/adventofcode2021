package main

import (
    "os"
    "bufio"
    "strings"
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
        ")": 3,
        "]": 57,
        "}": 1197,
        ">": 25137,
    }

    scanner := bufio.NewScanner(file)
    syntaxErrorScore := 0
    for scanner.Scan() {
        chunkStack := []string{}
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
                    score, _ := SCORE[char]
                    syntaxErrorScore += score
                    break
                }
                chunkStack = chunkStack[:len(chunkStack) - 1]
            }
        }
    }

    fmt.Println("Total syntax error score:", syntaxErrorScore)
}
