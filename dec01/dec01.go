package dec01

import (
    "os"
    "bufio"
    "strconv"
    "fmt"
)

func Helper(aggregate int) {
    file, _ := os.Open("01.txt")
    defer file.Close()

    depths := []int{}
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        depth, _ := strconv.Atoi(scanner.Text())
        depths = append(depths, depth)
    }

    numberOfIncreases := 0
    prevSlidingWindowValue := 0
	for i := 0; i < aggregate; i++ {
		prevSlidingWindowValue += depths[i]
	}
    index := aggregate
    for index < len(depths) {
        slidingWindowValue := 0

		for i := 0; i < aggregate; i++ {
			slidingWindowValue += depths[index - i]
		}

        if (slidingWindowValue > prevSlidingWindowValue) {
            numberOfIncreases += 1
        }
        prevSlidingWindowValue = slidingWindowValue
        index += 1
    }

    fmt.Println("number of increases: ", numberOfIncreases)
}
