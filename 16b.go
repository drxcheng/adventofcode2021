package main

import (
    // "os"
    // "bufio"
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

var versionNumberSum = 0

func getValue(values []int, operation int64) int {
    if operation == 0 {
        sum := 0
        for _, v := range values {
            sum += v
        }
        return sum
    } else if operation == 1 {
        product := 1
        for _, v := range values {
            product *= v
        }
        return product
    } else if operation == 2 {
        min := values[0]
        for i := 1; i < len(values); i++ {
            if values[i] < min {
                min = values[i]
            }
        }
        return min
    } else if operation == 3 {
        max := values[0]
        for i := 1; i < len(values); i++ {
            if values[i] > max {
                max = values[i]
            }
        }
        return max
    } else if operation == 5 {
        if values[0] > values[1] {
            return 1
        } else {
            return 0
        }
    } else if operation == 6 {
        if values[0] < values[1] {
            return 1
        } else {
            return 0
        }
    } else if operation == 7 {
        if values[0] == values[1] {
            return 1
        } else {
            return 0
        }
    } else {
        panic("wrong operation")
    }
}

func readPacket(bitString *[]rune, operation int64) int {
    if (len(*bitString) <= 6) {
        // fmt.Println("End of all packets")
        return 0
    }

    // fmt.Println("Read packet: ", string(bitString))
    version, _ := strconv.ParseInt(string((*bitString)[0:3]), 2, 64)
    *bitString = (*bitString)[3:]
    // fmt.Println("packet version:", version)
    versionNumberSum += int(version)
    // 3 bits T
    typeId, _ := strconv.ParseInt(string((*bitString)[0:3]), 2, 64)
    *bitString = (*bitString)[3:]
    // fmt.Println("typeId:", typeId)

    values := []int{}
    if typeId == 4 {
        // literal
        for len(*bitString) >= 5 {
            if (*bitString)[0] == '0' {
                literalNumber, _ := strconv.ParseInt(string((*bitString)[1:5]), 2, 64)
                values = append(values, int(literalNumber))
                // end of a packet
                allZero := true
                for _, char := range (*bitString)[5:] {
                    if char == '1' {
                        allZero = false
                        break
                    }
                }
                if !allZero {
                    subPacketString := (*bitString)[5:]
                    packetResult := readPacket(&subPacketString, -1)
                    values = append(values, packetResult)
                } else {
                    break
                }
            } else {
                literalNumber, _ := strconv.ParseInt(string((*bitString)[1:5]), 2, 64)
                values = append(values, int(literalNumber))
                *bitString = (*bitString)[5:]
            }
        }
    } else {
        // operation
        lengthTypeId := (*bitString)[0]
        *bitString = (*bitString)[1:]

        if lengthTypeId == '0' {
            // next 15
            length, _ := strconv.ParseInt(string((*bitString)[0:15]), 2, 64)
            // fmt.Println("length:", length)
            *bitString = (*bitString)[15:]

            subPacketString := (*bitString)[:length]
            packetResult := readPacket(&subPacketString, typeId)
            values = append(values, packetResult)

            if len(*bitString) > 0 {
                packetResult := readPacket(bitString, typeId)
                values = append(values, packetResult)
            }
        } else {
            // next 11
            count, _ := strconv.ParseInt(string((*bitString)[0:11]), 2, 64)
            // fmt.Println("count:", count)
            *bitString = (*bitString)[11:]
            for i := 0; i < int(count); i++ {
                packetResult := readPacket(bitString, typeId)
                values = append(values, packetResult)
            }
        }
    }

    return getValue(values, operation)
}

func main() {
    // file, _ := os.Open("16.txt")
    // defer file.Close()
    // scanner := bufio.NewScanner(file)
    // scanner.Scan()
    // hexString := scanner.Text()
    // hexString := "D2FE28"
    // hexString := "38006F45291200"
    // hexString := "EE00D40C823060"
    // hexString := "8A004A801A8002F478"
    // hexString := "620080001611562C8802118E34"
    // hexString := "C0015000016115A2E0802F182340"
    // hexString := "A0016C880162017C3686B18A3D4780"
    hexString := "C200B40A82" // [-8 1 2]
    // hexString := "04005AC33890" // [-7 6 9]
    // hexString := "880086C3E88112" // [-6 7 8 9]
    // hexString := "CE00C43D881120" // [-5 7 8 9]
    // hexString := "D8005AC2A8F0" //  [-2 5 15]
    // hexString := "F600BC2D8F" // [-3 5 15]
    // hexString := "9C005AC2F8F0" // [-1 5 15]
    // hexString := "9C0141080250320F1802104A08" // [-1 -8 1 3 -7 2 2]

    bitString := []rune{}
    for _, v := range strings.Split(hexString, "") {
        for _, char := range hexMapping[v] {
            bitString = append(bitString, char)
        }
    }
    // fmt.Println(string(bitString))

    result := readPacket(&bitString, -1)

    fmt.Println("versionNumberSum:", versionNumberSum)
    fmt.Println("result:", result)
}
