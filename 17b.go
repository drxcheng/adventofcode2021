package main

import (
    "fmt"
)

func canYBeInRange(velocity int, minSteps int, maxSteps int, min int, max int) bool {
    // fmt.Println("velocity:", velocity, "minSteps:", minSteps, "maxSteps:", maxSteps)
    v := velocity
    position := 0
    step := 0

    for ; step < minSteps; step++ {
        position += v
        v -= 1
    }

    // fmt.Println("position:", position)

    if position < min {
        // fmt.Println("position < min: ", position, min)
        return false
    } else if position >= min && position <= max {
        // fmt.Println("position >= min && position <= max: ", position, min, max)
        return true
    }

    for position >= max && (maxSteps == -1 || step < maxSteps) {
        position += v
        // fmt.Println("position:", position)
        if position >= min && position <= max {
            // fmt.Println("position >= min && position <= max: ", position, min, max)
            return true
        }
        v -= 1
        step += 1
    }

    // fmt.Println("other: ", position, min, max)
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
    fmt.Println(vXMin)

    valueCount := 0
    for vX := vXMin; vX <= targetXMax; vX++ {
        stepsMin := -1
        stepsMax := -1
        v := vX
        distance := 0
        for step := 1; step <= targetXMax; step++ {
            distance += v

            if stepsMin == -1 && distance >= targetXMin {
                stepsMin = step
            }

            if distance > targetXMax {
                stepsMax = step - 1
                break
            }

            if v > 0 {
                v -= 1
            }
        }

        if stepsMax > 0 && stepsMax < stepsMin {
            continue
        }
        // fmt.Println("x velocity=", vX, ", min steps:", stepsMin, ", max steps:", stepsMax)

        valueCountTmp := 0
        // This is so bad
        for vY := targetYMin; vY < 1000; vY++ {
            if canYBeInRange(vY, stepsMin, stepsMax, targetYMin, targetYMax) {
                valueCountTmp += 1
            }
        }

        // fmt.Println("tmp count: ", valueCountTmp)
        valueCount += valueCountTmp
    }

    fmt.Println("value count:", valueCount)
}
