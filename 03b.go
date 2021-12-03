package main

import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "fmt"
)

func filterDiagnosis(listOfDiagnosis [][]rune, position int) ([][]rune, [][]rune) {
    newListOfDiagnosisHigh := [][]rune{}
    newListOfDiagnosisLow := [][]rune{}
    for _, diagnosis := range listOfDiagnosis {
        if diagnosis[position] == '1' {
            newListOfDiagnosisHigh = append(newListOfDiagnosisHigh, diagnosis)
        } else {
            newListOfDiagnosisLow = append(newListOfDiagnosisLow, diagnosis)
        }
    }

    return newListOfDiagnosisHigh, newListOfDiagnosisLow
}

func main() {
    file, _ := os.Open("03.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    numberOfLines := 0
    listOfDiagnosis := [][]rune{}
    for scanner.Scan() {
        diagnosis := []rune(scanner.Text())
        listOfDiagnosis = append(listOfDiagnosis, diagnosis)
        numberOfLines += 1
    }

    lengthOfDiagnosis := len(listOfDiagnosis[0])

    oxygenDiagnosisCandidates := listOfDiagnosis
    co2DiagnosisCandidates := listOfDiagnosis
    for i := 0; i < lengthOfDiagnosis; i++ {
        if len(oxygenDiagnosisCandidates) > 1 {
            listOfDiagnosisHigh, listOfDiagnosisLow := filterDiagnosis(oxygenDiagnosisCandidates, i)
            if len(listOfDiagnosisHigh) >= len(listOfDiagnosisLow) {
                oxygenDiagnosisCandidates = listOfDiagnosisHigh
            } else {
                oxygenDiagnosisCandidates = listOfDiagnosisLow
            }
        }

        if len(co2DiagnosisCandidates) > 1 {
            listOfDiagnosisHigh, listOfDiagnosisLow := filterDiagnosis(co2DiagnosisCandidates, i)
            if len(listOfDiagnosisLow) <= len(listOfDiagnosisHigh) {
                co2DiagnosisCandidates = listOfDiagnosisLow
            } else {
                co2DiagnosisCandidates = listOfDiagnosisHigh
            }
        }
    }

    oxygenRateBits := []string{}
    co2RateBits := []string{}
    for i := 0; i < lengthOfDiagnosis; i++ {
        oxygenRateBits = append(oxygenRateBits, string(oxygenDiagnosisCandidates[0][i]))
        co2RateBits = append(co2RateBits, string(co2DiagnosisCandidates[0][i]))
    }

    oxygenRate, _ := strconv.ParseInt(strings.Join(oxygenRateBits[:], ""), 2, 64)
    co2Rate, _ := strconv.ParseInt(strings.Join(co2RateBits[:], ""), 2, 64)

    fmt.Println("oxygen rate: ", oxygenRate)
    fmt.Println("co2 rate: ", co2Rate)
    fmt.Println("life support rating: ", oxygenRate * co2Rate)
}
