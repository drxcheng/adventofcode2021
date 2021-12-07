package dec05

import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "fmt"
)

// for each fish with original timer being t (t is between 1 and 5)
// calculate how many fishes there is after N days
/*
 0d: 1       2       3       4       5
 7d: 1,3     2,4     3,5     4,6     5,7
 8d: 0,2     1,3     2,4     3,4     4,6  (reminder = 1)
 9d: 6,1,8   0,2     1,3     2,4     3,5  (reminder = 2)
10d: 5,0,7   6,1,8   0,2     1,3     2,4  (reminder = 3)
11d: 4,6,6,8 5,0,7   6,1,8   0,2     1,3  (reminder = 4)
14d: 1,3,3,5 2,4,4,6 3,5,5,7 4,6,6,8 5,7,0
*/

func calcFishNumber(initialTimer int, numberOfCycles int, remainderDays int) int {
    timerCountMap := make(map[int]int, 9)
    for i := 0; i < 9; i++ {
        timerCountMap[i] = 0
    }
    timerCountMap[initialTimer] = 1

    for cycle := 0; cycle < numberOfCycles; cycle++ {
        // 0 -> 0,2
        // 1 -> 1,3
        // 2 -> 2,4
        // 3 -> 3,5
        // 4 -> 4,6
        // 5 -> 5,7
        // 6 -> 6,8
        // 7 -> 0
        // 8 -> 1
        timerCountMapNewRound := make(map[int]int, 9)
        for i := 0; i < 9; i++ {
            timerCountMapNewRound[i] = 0
        }
        for i, timerCount := range timerCountMap {
            if i < 7 {
                timerCountMapNewRound[i] += timerCount
                timerCountMapNewRound[i + 2] += timerCount
            } else {
                timerCountMapNewRound[i - 7] += timerCount
            }
        }
        for i, timerCount := range timerCountMapNewRound {
            timerCountMap[i] = timerCount
        }
        // fmt.Println(timerCountMap)
    }

    totalCount := 0
    for _, timerCount := range timerCountMap {
        totalCount += timerCount
    }

    i := 0
    for i < remainderDays {
        totalCount += timerCountMap[i]
        i += 1
    }

    return totalCount
}

func Helper(DAY int) {
    file, _ := os.Open("06.txt")
    defer file.Close()

    numberOfCycles := int(DAY / 7)
    remainderDays := DAY - numberOfCycles * 7

    fmt.Println("cycles:", numberOfCycles, ", remainder days:", remainderDays)

    timerMap := make(map[int]int, 7)
    initalCountMap := make(map[int]int, 7)
    for i := 1; i < 7; i++ {
        numberOfFishes := calcFishNumber(i, numberOfCycles, remainderDays)
        fmt.Println("initial timer:", i, ", # of fishes:", numberOfFishes)
        timerMap[i] = numberOfFishes
        initalCountMap[i] = 0
    }

    // fmt.Println(timerMap)

    scanner := bufio.NewScanner(file)
    scanner.Scan()
    for _, v := range strings.Split(scanner.Text(), ",") {
        timer, _ := strconv.Atoi(v)
        initalCountMap[timer] += 1
    }

    // fmt.Println(initalCountMap)
    totalCount := 0
    for i := 0; i < len(initalCountMap); i++ {
        totalCount += initalCountMap[i] * timerMap[i]
    }

    fmt.Println("Total number of fishes: ", totalCount)
}
