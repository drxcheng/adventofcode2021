package main

import (
    "os"
    "bufio"
    "strconv"
    "fmt"
    "strings"
)

func main() {
    file, _ := os.Open("03.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    numberOfLines := 0
    sumOfEachPosition := []int{}
    for scanner.Scan() {
        diagnosis := []rune(scanner.Text())
        for i := 0; i < len(diagnosis); i++ {
            bit, _ := strconv.Atoi(string(diagnosis[i]))
            if bit != 0 && bit != 1 {
                os.Exit(bit)
            }

            if len(sumOfEachPosition) <= i {
                sumOfEachPosition = append(sumOfEachPosition, 0)
            }

            sumOfEachPosition[i] += bit // 0 or 1
        }
        numberOfLines += 1
    }

    fmt.Println("Total lines: ", numberOfLines)

    gammaRateBits := []string{}
    epsilonRateBits := []string{}
    for i := 0; i < len(sumOfEachPosition); i++ {
        gammaRateBitInt := 0
        if sumOfEachPosition[i] > numberOfLines / 2 {
            gammaRateBitInt = 1
        }
        gammaRateBit := strconv.Itoa(gammaRateBitInt)
        epsilonRateBit := strconv.Itoa(1 - gammaRateBitInt)
        gammaRateBits = append(gammaRateBits, gammaRateBit)
        epsilonRateBits = append(epsilonRateBits, epsilonRateBit)
    }

    gammaRateString := strings.Join(gammaRateBits[:], "")
    epsilonRateString := strings.Join(epsilonRateBits[:], "")

    gammaRate, _ := strconv.ParseInt(gammaRateString, 2, 64)
    epsilonRate, _ := strconv.ParseInt(epsilonRateString, 2, 64)

    fmt.Println("gamma rate: ", gammaRate)
    fmt.Println("epsilon rate: ", epsilonRate)
    fmt.Println("power consumption: ", gammaRate * epsilonRate)
}
