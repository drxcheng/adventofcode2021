package main

import (
    dec07 "./dec07"
)

func calcFuelCost(steps int) int {
    return (steps + 1) * steps / 2
}

// run GO111MODULE=off go run 07b.go
func main() {
    dec07.Helper(calcFuelCost)
}
