package dec07

import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "fmt"
    "math"
)

type calcFuelCostType func(int) int

func Helper(calcFuelCost calcFuelCostType) {
    file, _ := os.Open("07.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Scan()
    positions := []int{}
    maxPosition := 0
    for _, v := range strings.Split(scanner.Text(), ",") {
        position, _ := strconv.Atoi(v)
        positions = append(positions, position)
        if position > maxPosition {
            maxPosition = position
        }
    }

    // fmt.Println(maxPosition)

    leastFuel := calcFuelCost(maxPosition) * len(positions)
    leastFuelPosition := -1
    for i := 0; i <= maxPosition; i++ {
        fuel := 0
        for _, p := range positions {
            fuel += calcFuelCost(int(math.Abs(float64(p - i))))
        }

        if fuel < leastFuel {
            leastFuel = fuel
            leastFuelPosition = i
        }
    }

    fmt.Println("Position:", leastFuelPosition)
    fmt.Println("Fuel:", leastFuel)
}
