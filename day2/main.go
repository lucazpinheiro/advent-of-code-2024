package main

import (
	"bufio"
	"fmt"
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

			if shouldIncrease {
				if level > next || level == next {
					break
				}
				if next-level > 3 {
					break
				}
			} else {
				if level < next || level == next {
					break
				}
				if level-next > 3 {
					break
				}
			}

			if i+1 == len(y)-1 {
				safeCount++
			}
		}
	}

	fmt.Println("day 2 part one:", safeCount)
}

func partTwo() {
	readFile, err := os.Open("sample.txt")
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
		tries := 0
		x := fileScanner.Text()
		y := strings.Split(x, " ")

		safe := false
		failedAt := 0

		for tries < 2 {
			tries++
			safe, failedAt = findFailLevel(y)
			fmt.Println(safe, failedAt)
			if !safe {
				y = append(y[:failedAt], y[failedAt+1:]...)
			}
		}
		// fmt.Println(safe, failedAt)
		if safe {
			safeCount++
		}
	}

	fmt.Println("day 2 part two", safeCount)
}

func findFailLevel(y []string) (bool, int) {
	safeCount := 0
	failedAt := 0

	first := utils.SafeInt(y[0])
	second := utils.SafeInt(y[1])

	shouldIncrease := first < second
	for i := 0; i < len(y)-1; i++ {
		level := utils.SafeInt(y[i])
		next := utils.SafeInt(y[i+1])

		if shouldIncrease {
			if level > next || level == next {
				failedAt = i
				break
			}
			if next-level > 3 {
				failedAt = i
				break
			}
		} else {
			if level < next || level == next {
				failedAt = i + 1
				break
			}
			if level-next > 3 {
				failedAt = i + 1
				break
			}
		}

		if i+1 == len(y)-1 {
			safeCount++
		}
	}

	if safeCount > 0 {
		return true, 0
	} else {
		return false, failedAt
	}

}
