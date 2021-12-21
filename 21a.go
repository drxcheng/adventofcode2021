package main

import (
    "fmt"
)

func rollThreeTimes(pos *int, diceValue *int, numberOfRolls *int) {
    steps := *diceValue * 3 + 3
    *diceValue += 3
    *numberOfRolls += 3

    *pos = (*pos + steps - 1) % 10 + 1
}

const WINNING_SCORE = 1000

func main() {
    player1Pos := 10
    player2Pos := 8
    diceValue := 1
    numberOfRolls := 0
    player1Score := 0
    player2Score := 0

    var losingPlayerScore int
    for {
        // player 1
        rollThreeTimes(&player1Pos, &diceValue, &numberOfRolls)
        player1Score += player1Pos
        if player1Score >= WINNING_SCORE {
            // game over
            losingPlayerScore = player2Score
            break
        }
        // fmt.Println("player 1 score: ", player1Score)

        // player 2
        rollThreeTimes(&player2Pos, &diceValue, &numberOfRolls)
        player2Score += player2Pos
        if player2Score >= WINNING_SCORE {
            // game over
            losingPlayerScore = player1Score
            break
        }

        // fmt.Println("player 2 score: ", player2Score)
    }

    fmt.Println("losingPlayerScore: ", losingPlayerScore)
    fmt.Println("numberOfRolls: ", numberOfRolls)
    fmt.Println("result: ", losingPlayerScore * numberOfRolls)
}
