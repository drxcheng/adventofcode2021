package dec12

import (
    "os"
    "bufio"
    "strings"
    "fmt"
)

const START = "start"
const END = "end"

func isSmallCave(cave string) bool {
    return cave != START && cave != END && strings.ToLower(cave) == cave
}

func canVisitSmallCave(currentPath []string, nextCave string, allowRevisit bool) bool {
    if allowRevisit {
        smallCaveVisitedCount := map[string]int {nextCave: 1}

        for _, cave := range currentPath {
            if isSmallCave(cave) {
                _, ok := smallCaveVisitedCount[cave]
                if !ok {
                    smallCaveVisitedCount[cave] = 0
                }
                smallCaveVisitedCount[cave] += 1
            }
        }

        hasMultipleVisit := false
        for _, count := range smallCaveVisitedCount {
            if count > 2 {
                return false
            } else if count > 1 {
                if hasMultipleVisit {
                    return false
                }

                hasMultipleVisit = true
            }
        }

        return true
    }

    for _, cave := range currentPath {
        if cave == nextCave {
            return false
        }
    }

    return true
}

func recursivelyTraverse(caveMap map[string][]string, allPaths [][]string, currentPath []string, from string, allowRevisit bool) [][]string {
    for _, nextCave := range caveMap[from] {
        // fmt.Println("Current path:", currentPath, ", next cave:", nextCave)
        newPath := []string{}
        for _, cave := range currentPath {
            newPath = append(newPath, cave)
        }
        newPath = append(newPath, nextCave)

        if nextCave == END {
            // fmt.Println("to the end. Final path: ", newPath)
            allPaths = append(allPaths, newPath)
        } else if !isSmallCave(nextCave) {
            // fmt.Println("Continue. New path: ", newPath)
            allPaths = recursivelyTraverse(caveMap, allPaths, newPath, nextCave, allowRevisit)
        } else if canVisitSmallCave(currentPath, nextCave, allowRevisit) {
            // fmt.Println("Continue. New path: ", newPath)
            allPaths = recursivelyTraverse(caveMap, allPaths, newPath, nextCave, allowRevisit)
        }
    }

    return allPaths
}

func Helper(allowRevisit bool) {
    file, _ := os.Open("12.txt")
    defer file.Close()

    caveMap := make(map[string][]string)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        caves := strings.Split(scanner.Text(), "-")
        if len(caves) != 2 {
            panic("Invalid line")
        }

        if caves[0] != END {
            _, ok := caveMap[caves[0]]
            if !ok {
                caveMap[caves[0]] = []string{}
            }

            if caves[1] != START {
                caveMap[caves[0]] = append(caveMap[caves[0]], caves[1])
            }
        }

        if caves[1] != END {
            _, ok := caveMap[caves[1]]
            if !ok {
                caveMap[caves[1]] = []string{}
            }

            if caves[0] != START {
                caveMap[caves[1]] = append(caveMap[caves[1]], caves[0])
            }
        }
    }
    // fmt.Println(caveMap)

    allPaths := [][]string{}
    currentPath := []string{START}
    allPaths = recursivelyTraverse(caveMap, allPaths, currentPath, START, allowRevisit)

    // for i, path := range allPaths {
    //     fmt.Println((i + 1), ":", path)
    // }

    fmt.Println("Number of paths:", len(allPaths))
}
