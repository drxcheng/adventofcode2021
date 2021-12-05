package dec04

import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "fmt"
)

const SIZE int = 5

func readOneBoard(scanner *bufio.Scanner) [SIZE * SIZE]int {
    var board [SIZE * SIZE]int
    for m := 0; m < SIZE; m++ {
        scanner.Scan()
        numbers := strings.Fields(scanner.Text())
        for n := 0; n < SIZE; n++ {
            board[m * SIZE + n], _ = strconv.Atoi(numbers[n])
        }
    }

    return board
}

func calcScore(board [SIZE * SIZE]int, sequence []int) (int, int) {
    // make a map first
    boardMap := make(map[int]int, SIZE * SIZE)
    for i, number := range board {
        boardMap[number] = i
    }
    for index, currentNumber := range sequence {
        if index, ok := boardMap[currentNumber]; ok {
            board[index] = -1
        }

        if doesBoardWin(board) {
            unmarkedSum := 0
            for _, number := range board {
                if number != -1 {
                    unmarkedSum += number
                }
            }

            return index, unmarkedSum * currentNumber;
        }
    }

    return len(sequence), 0
}

func doesBoardWin(board [SIZE * SIZE]int) bool {
    for m := 0; m < SIZE; m++ {
        sum := 0
        for n := 0; n < SIZE; n++ {
            sum += board[m * SIZE + n]
        }
        if sum == -SIZE {
            return true
        }
    }

    for n := 0; n < SIZE; n++ {
        sum := 0
        for m := 0; m < SIZE; m++ {
            sum += board[m * SIZE + n]
        }
        if sum == -SIZE {
            return true
        }
    }

    return false
}

func Helper(needToWin bool) {
    file, _ := os.Open("04.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Scan()
    sequence := []int{}
    for _, v := range strings.Split(scanner.Text(), ",") {
        number, _ := strconv.Atoi(v)
        sequence = append(sequence, number)
    }

    leastNumberOfSteps := len(sequence)
    mostNumberOfSteps := SIZE
    boardScore := 0
    for scanner.Scan() {
        // an empty line first
        board := readOneBoard(scanner)
        numberOfSteps, score := calcScore(board, sequence)

        if needToWin && numberOfSteps < leastNumberOfSteps {
            leastNumberOfSteps = numberOfSteps
            boardScore = score
        } else if !needToWin && numberOfSteps > mostNumberOfSteps {
            mostNumberOfSteps = numberOfSteps
            boardScore = score
        }
    }

    fmt.Println("board score: ", boardScore)
}
