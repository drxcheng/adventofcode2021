package dec02

import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "math"
    "fmt"
)

func Helper(useAim bool) {
    file, _ := os.Open("02.txt")
    defer file.Close()

    offsetHorizontal := 0
    offsetVertical := 0
    aim := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lineData := strings.Split(scanner.Text(), " ")
        direction := lineData[0]
        amount, _ := strconv.Atoi(lineData[1])

        if (direction == "forward") {
            offsetHorizontal += amount
            if useAim {
                offsetVertical += amount * aim
            }
        } else if (direction == "down") {
            if useAim {
                aim += amount
            } else {
                offsetVertical += amount
            }
        } else if (direction == "up") {
            if useAim {
                aim -= amount
            } else {
                offsetVertical -= amount
            }
        } else {
            panic("Wrong direction")
        }
    }

    // fmt.Println("horizontal position: ", offsetHorizontal)
    // fmt.Println("vertical position: ", offsetVertical)
    fmt.Println("result: ", int(math.Abs(float64(offsetHorizontal * offsetVertical))))
}
