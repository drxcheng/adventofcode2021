package main

import (
    "fmt"
)

func getHighestPosition(velocity int) int {
    v := velocity
    position := 0
    for step := 0; v > 0; step++ {
        position += v
        v -= 1
    }

    return position
}

func canYBeInRange(velocity int, minSteps int, min int, max int) bool {
    v := velocity
    position := 0
    for step := 0; step < minSteps; step++ {
        position += v
        v -= 1
    }

    if position < min {
        return false
    }

    for position > max {
        position += v
        if position >= min && position <= max {
            return true
        }
        v -= 1
    }

    return false
}

func main() {
    // sample: target area: x=20..30, y=-10..-5
    // targetXMin := 20
    // targetXMax := 30
    // targetYMin := -10
    // targetYMax := -5
    // real: target area: x=94..151, y=-156..-103
    targetXMin := 94
    targetXMax := 151
    targetYMin := -156
    targetYMax := -103

    vXMin := 0
    for vX := 1; vX < targetXMin; vX++ {
        maxX := vX * (vX + 1) / 2;
        if maxX > targetXMin {
            vXMin = vX
            break
        }
    }

    stepsMin := -1
    v := vXMin
    distance := 0
    for step := 1; step < targetXMax; step++ {
        distance += v

        if stepsMin == -1 && distance >= targetXMin {
            stepsMin = step
        }

        if distance > targetXMax {
            break
        }

        if v > 0 {
            v -= 1
        }
    }

    fmt.Println("x velocity=", vXMin, ", min steps:", stepsMin)

    highestPosition := 0
    // this works but is BAD
    for vY := stepsMin; vY < 1000; vY++ {
        // fmt.Println("y velocity=", vY)
        if canYBeInRange(vY, stepsMin, targetYMin, targetYMax) {
            // fmt.Println("can be in range")
            position := getHighestPosition(vY)
            if position > highestPosition {
                highestPosition = position
                // fmt.Println("y velocity=", vY, ", highest position:", highestPosition)
            }
        } else {
            // fmt.Println("cannot be in range")
            // break
        }
    }

    fmt.Println("highest position:", highestPosition)
}
