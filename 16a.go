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

func readPacket(bitString []rune, versionNumberSum *int) []rune {
    // fmt.Println("Read packet")
    version, _ := strconv.ParseInt(string(bitString[0:3]), 2, 64)
    bitString = bitString[3:]
    fmt.Println("packet version:", version)
    *versionNumberSum += int(version)
    // 3 bits T
    typeId, _ := strconv.ParseInt(string(bitString[0:3]), 2, 64)
    bitString = bitString[3:]
    // fmt.Println(typeId)
    if typeId == 4 {
        // literal
        for {
            if bitString[0] == '0' {
                // end of a packet
                if len(bitString) < 5 {
                    fmt.Println(bitString)
                    return bitString[len(bitString) - 1:]
                } else {
                    return bitString[5:]
                }
            }
            bitString = bitString[5:]
        }
    } else {
        // operation
        lengthTypeId := bitString[0]
        bitString = bitString[1:]

        if lengthTypeId == '0' {
            // next 15
            length, _ := strconv.ParseInt(string(bitString[0:15]), 2, 64)
            fmt.Println("length:", length)
            bitString = bitString[15:]
            readPacket(bitString[:length], versionNumberSum)
            bitString = bitString[length:]
        } else {
            // next 11
            count, _ := strconv.ParseInt(string(bitString[0:11]), 2, 64)
            fmt.Println("count:", count)
            bitString = bitString[11:]
            for i := 0; i < int(count); i++ {
                bitString = readPacket(bitString, versionNumberSum)
            }
        }

        return bitString
    }
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
    // hexString := "620080001611562C8802118E34" // expect 12, actual 7
    // hexString := "C0015000016115A2E0802F182340" // expect 23, actual 6
    // hexString := "A0016C880162017C3686B18A3D4780"
    hexString := "2056FA18025A00A4F52AB13FAB6CDA779E1B2012DB003301006A35C7D882200C43289F07A5A192D200C1BC011969BA4A485E63D8FE4CC80480C00D500010F8991E23A8803104A3C425967260020E551DC01D98B5FEF33D5C044C0928053296CDAFCB8D4BDAA611F256DE7B945220080244BE59EE7D0A5D0E6545C0268A7126564732552F003194400B10031C00C002819C00B50034400A70039C009401A114009201500C00B00100D00354300254008200609000D39BB5868C01E9A649C5D9C4A8CC6016CC9B4229F3399629A0C3005E797A5040C016A00DD40010B8E508615000213112294749B8D67EC45F63A980233D8BCF1DC44FAC017914993D42C9000282CB9D4A776233B4BF361F2F9F6659CE5764EB9A3E9007ED3B7B6896C0159F9D1EE76B3FFEF4B8FCF3B88019316E51DA181802B400A8CFCC127E60935D7B10078C01F8B50B20E1803D1FA21C6F300661AC678946008C918E002A72A0F27D82DB802B239A63BAEEA9C6395D98A001A9234EA620026D1AE5CA60A900A4B335A4F815C01A800021B1AE2E4441006A0A47686AE01449CB5534929FF567B9587C6A214C6212ACBF53F9A8E7D3CFF0B136FD061401091719BC5330E5474000D887B24162013CC7EDDCDD8E5E77E53AF128B1276D0F980292DA0CD004A7798EEEC672A7A6008C953F8BD7F781ED00395317AF0726E3402100625F3D9CB18B546E2FC9C65D1C20020E4C36460392F7683004A77DB3DB00527B5A85E06F253442014A00010A8F9106108002190B61E4750004262BC7587E801674EB0CCF1025716A054AD47080467A00B864AD2D4B193E92B4B52C64F27BFB05200C165A38DDF8D5A009C9C2463030802879EB55AB8010396069C413005FC01098EDD0A63B742852402B74DF7FDFE8368037700043E2FC2C8CA00087C518990C0C015C00542726C13936392A4633D8F1802532E5801E84FDF34FCA1487D367EF9A7E50A43E90"

    bitString := []rune{}
    for _, v := range strings.Split(hexString, "") {
        for _, char := range hexMapping[v] {
            bitString = append(bitString, char)
        }
    }
    fmt.Println(string(bitString))

    versionNumberSum := 0
    readPacket(bitString, &versionNumberSum)

    fmt.Println("versionNumberSum:", versionNumberSum)
}
