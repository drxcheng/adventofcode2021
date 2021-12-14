package dec14

import (
    "os"
    "bufio"
    "strings"
    "sort"
    "fmt"
)

func Helper(STEP int) {
    file, _ := os.Open("14.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Scan()
    template := scanner.Text()
    scanner.Scan()

    insertionRulesMap := make(map[string]string)
    for scanner.Scan() {
        insertionLineParts := strings.Split(scanner.Text(), " -> ")
        insertionRulesMap[insertionLineParts[0]] = insertionLineParts[1]
    }
    // fmt.Println("insertionRulesMap:", insertionRulesMap)

    pairCountMap := make(map[string]int)
    for i := 0; i <= len(template) - 2; i++ {
        pairCountMap[template[i:i+2]] += 1
    }
    // fmt.Println(pairCountMap)

    for step := 0; step < STEP; step++ {
        newPairCountMap := make(map[string]int)
        for pair, count := range pairCountMap {
            insertion, found := insertionRulesMap[pair]
            if found {
                newPairCountMap[pair[:1] + insertion] += count
                newPairCountMap[insertion + pair[1:]] += count
            }
        }

        // fmt.Println(newPairCountMap)
        pairCountMap = newPairCountMap
    }

    // count
    charCountMap := make(map[string]int)
    // first record start and end of template
    charCountMap[template[:1]] = 1
    charCountMap[template[len(template) - 1:]] = 1
    for pair, count := range pairCountMap {
        charCountMap[pair[:1]] += count
        charCountMap[pair[1:]] += count
    }
    // fmt.Println(charCountMap)

    sortedCount := []int{}
    for _, count := range charCountMap {
        sortedCount = append(sortedCount, count/2)
    }
    sort.Ints(sortedCount)
    // fmt.Println(sortedCount)

    fmt.Println("result =", sortedCount[len(sortedCount) - 1] - sortedCount[0])
}
