package main

import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "fmt"
)

var hexMapping = map[string]string{
    "0": "0000",
    "1": "0001",
    "2": "0010",
    "3": "0011",
    "4": "0100",
    "5": "0101",
    "6": "0110",
    "7": "0111",
    "8": "1000",
    "9": "1001",
    "A": "1010",
    "B": "1011",
    "C": "1100",
    "D": "1101",
    "E": "1110",
    "F": "1111",
}

func readPacket(bitString []rune, versionNumberSum *int) []rune {
    if (len(bitString) <= 6) {
        // fmt.Println("End of all packets")
        return []rune{}
    }

    // fmt.Println("Read packet: ", string(bitString))
    version, _ := strconv.ParseInt(string(bitString[0:3]), 2, 64)
    bitString = bitString[3:]
    // fmt.Println("packet version:", version)
    *versionNumberSum += int(version)
    // 3 bits T
    typeId, _ := strconv.ParseInt(string(bitString[0:3]), 2, 64)
    bitString = bitString[3:]
    // fmt.Println(typeId)
    if typeId == 4 {
        // literal
        for len(bitString) >= 5 {
            if bitString[0] == '0' {
                // end of a packet
                allZero := true
                for _, char := range bitString[5:] {
                    if char == '1' {
                        allZero = false
                        break
                    }
                }
                if !allZero {
                    bitString = readPacket(bitString[5:], versionNumberSum)
                } else {
                    return []rune{}
                }
            } else {
                bitString = bitString[5:]
            }
        }
        return bitString
    } else {
        // operation
        lengthTypeId := bitString[0]
        bitString = bitString[1:]

        if lengthTypeId == '0' {
            // next 15
            length, _ := strconv.ParseInt(string(bitString[0:15]), 2, 64)
            // fmt.Println("length:", length)
            bitString = bitString[15:]

            readPacket(bitString[:length], versionNumberSum)
            bitString = bitString[length:]

            if len(bitString) > 0 {
                bitString = readPacket(bitString, versionNumberSum)
            }
        } else {
            // next 11
            count, _ := strconv.ParseInt(string(bitString[0:11]), 2, 64)
            // fmt.Println("count:", count)
            bitString = bitString[11:]
            for i := 0; i < int(count); i++ {
                bitString = readPacket(bitString, versionNumberSum)
            }
        }

        return bitString
    }
}

func main() {
    file, _ := os.Open("16.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)
    scanner.Scan()
    hexString := scanner.Text()
    // hexString := "D2FE28"
    // hexString := "38006F45291200"
    // hexString := "EE00D40C823060"
    // hexString := "8A004A801A8002F478"
    // hexString := "620080001611562C8802118E34"
    // hexString := "C0015000016115A2E0802F182340"
    // hexString := "A0016C880162017C3686B18A3D4780"

    bitString := []rune{}
    for _, v := range strings.Split(hexString, "") {
        for _, char := range hexMapping[v] {
            bitString = append(bitString, char)
        }
    }
    // fmt.Println(string(bitString))

    versionNumberSum := 0
    bitString = readPacket(bitString, &versionNumberSum)

    fmt.Println("versionNumberSum:", versionNumberSum)
}
