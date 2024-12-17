package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	utils "github.com/lucazpinheiro/advent-of-code-2024"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		readFile.Close()
	}()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	safeCount := 0

	for fileScanner.Scan() {
		x := fileScanner.Text()
		y := strings.Split(x, " ")

		first := utils.SafeInt(y[0])
		second := utils.SafeInt(y[1])

		shouldIncrease := first < second

		for i := 0; i < len(y)-1; i++ {
			level := utils.SafeInt(y[i])
			next := utils.SafeInt(y[i+1])

			diff := int(math.Abs(float64(level - next)))

			if diff == 0 {
				break
			}

			if diff > 3 {
				break
			}

			if shouldIncrease && level > next {
				break
			}

			if !shouldIncrease && level < next {
				break
			}

			if i+1 == len(y)-1 {
				safeCount++
			}
		}
	}

	fmt.Println("day 2 part one:", safeCount)
}

func partTwo() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		readFile.Close()
	}()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	safeCount := 0

	for fileScanner.Scan() {
		x := fileScanner.Text()
		y := strings.Split(x, " ")

		if checkSafetyWithDeletion(utils.StringToInt(y)) {
			safeCount++
		}
	}

	fmt.Println("day 2 part one:", safeCount)
}

func checkSafetyWithDeletion(reportNum []int) bool {
	fmt.Println(reportNum)
	safe := false
	skip := 0 // Start by skipping the first element
	for i := 0; i < len(reportNum); i++ {
		// Create a new slice excluding the skipped element
		iterationSlice := []int{}
		for j := 0; j < len(reportNum); j++ {
			if j == skip {
				continue // Skip the current index
			}
			iterationSlice = append(iterationSlice, reportNum[j])
		}
		fmt.Println(iterationSlice)
		if isSafe(iterationSlice) {
			safe = true
		}

		// Update the skip index for the next iteration
		skip = (skip + 1) % len(reportNum)
	}

	return safe
}

func isSafe(list []int) bool {
	safe := true
	first := list[0]
	second := list[1]

	shouldIncrease := first < second

	for i := 0; i < len(list)-1; i++ {
		level := list[i]
		next := list[i+1]

		diff := int(math.Abs(float64(level - next)))

		if diff == 0 {
			fmt.Println(1)
			safe = false
			break
		}

		if diff > 3 {
			fmt.Println(2)
			safe = false
			break
		}

		if shouldIncrease && level > next {
			fmt.Println(3)
			safe = false
			break
		}

		if !shouldIncrease && level < next {
			fmt.Println(4)
			safe = false
			break
		}
	}

	return safe
}
