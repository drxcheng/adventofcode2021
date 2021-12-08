package main

import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "sort"
    "fmt"
)

/*
    aaaa
   b    c
   b    c
    dddd
   e    f
   e    f
    gggg
 */
// from 1 (2 segments) and 7 (3 segments), we can derive a, also (c, f) (don't know which one yet)
// from 4 (4 segments) and 1 (2 segments), we can derive (b, d) (don't know which one yet)
// from 0 (6 segments), 6 (6 segments), 9 (6 segments), we can derive all

func SortString(digit string) string {
    segments := strings.Split(digit, "")
    sort.Strings(segments)
    return strings.Join(segments, "")
}

// variable naming is a mess
func main() {
    file, _ := os.Open("08.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    totalValue := 0
    for scanner.Scan() {
        data := strings.Split(scanner.Text(), " | ")
        var digits [10]string
        segmentsToDigit := make(map[string]string)
        digit1 := ""
        digit4 := ""
        digit7 := ""
        for i, digit := range strings.Fields(data[0]) {
            sortedDigit := SortString(digit)
            digits[i] = sortedDigit

            if len(sortedDigit) == 2 {
                segmentsToDigit[sortedDigit] = "1"
                digit1 = sortedDigit
            } else if len(sortedDigit) == 4 {
                segmentsToDigit[sortedDigit] = "4"
                digit4 = sortedDigit
            } else if len(sortedDigit) == 3 {
                segmentsToDigit[sortedDigit] = "7"
                digit7 = sortedDigit
            } else if len(sortedDigit) == 7 {
                segmentsToDigit[sortedDigit] = "8"
            }
        }

        segmentA := ""
        segmentB := ""
        segmentC := ""
        segmentD := ""
        segmentE := ""
        segmentF := ""
        segmentG := ""
        segmentCF := ""
        segmentBD := ""
        segmentEG := "abcdefg"
        for _, segment := range strings.Split(digit7, "") {
            if strings.Contains(digit1, segment) {
                segmentCF += segment
            } else {
                segmentA = segment
            }

            segmentEG = strings.Replace(segmentEG, segment, "", 1)
        }
        for _, segment := range strings.Split(digit4, "") {
            if !strings.Contains(digit1, segment) {
                segmentBD += segment
            }

            segmentEG = strings.Replace(segmentEG, segment, "", 1)
        }

        cfCount := 0
        bdCount := 0
        egCount := 0
        for _, segments := range digits {
            if len(segments) == 6 {
                if strings.Contains(segments, segmentCF[0:1]) {
                    cfCount += 1
                }
                if strings.Contains(segments, segmentBD[0:1]) {
                    bdCount += 1
                }
                if strings.Contains(segments, segmentEG[0:1]) {
                    egCount += 1
                }
            }
        }
        if cfCount == 3 {
            segmentF = segmentCF[0:1]
            segmentC = segmentCF[1:2]
        } else {
            segmentC = segmentCF[0:1]
            segmentF = segmentCF[1:2]
        }
        if bdCount == 3 {
            segmentB = segmentBD[0:1]
            segmentD = segmentBD[1:2]
        } else {
            segmentD = segmentBD[0:1]
            segmentB = segmentBD[1:2]
        }
        if egCount == 3 {
            segmentG = segmentEG[0:1]
            segmentE = segmentEG[1:2]
        } else {
            segmentE = segmentEG[0:1]
            segmentG = segmentEG[1:2]
        }

        segmentsToDigit[SortString(segmentA + segmentB + segmentC + segmentE + segmentF + segmentG)] = "0"
        segmentsToDigit[SortString(segmentA + segmentC + segmentD + segmentE + segmentG)] = "2"
        segmentsToDigit[SortString(segmentA + segmentC + segmentD + segmentF + segmentG)] = "3"
        segmentsToDigit[SortString(segmentA + segmentB + segmentD + segmentF + segmentG)] = "5"
        segmentsToDigit[SortString(segmentA + segmentB + segmentD + segmentE + segmentF + segmentG)] = "6"
        segmentsToDigit[SortString(segmentA + segmentB + segmentC + segmentD + segmentF + segmentG)] = "9"

        // fmt.Println(segmentsToDigit)

        valueString := ""
        for _, digit := range strings.Fields(data[1]) {
            sortedDigit := SortString(digit)
            valueString += segmentsToDigit[sortedDigit]
        }
        // fmt.Println(valueString)
        value, _ := strconv.Atoi(valueString)
        totalValue += value
    }

    fmt.Println("Total value:", totalValue)
}
