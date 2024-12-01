package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

	var leftList, rightList []int

	for fileScanner.Scan() {
		x := fileScanner.Text()
		y := strings.Split(x, " ")

		leftList = append(leftList, utils.SafeInt(y[0]))
		rightList = append(rightList, utils.SafeInt(y[len(y)-1]))
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	totalDistance := 0

	for i := 0; i < len(leftList); i++ {
		d := leftList[i] - rightList[i]
		if d < 0 {
			totalDistance += -d
		} else {
			totalDistance += d
		}
	}

	fmt.Println("day 1 part one:", totalDistance)
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

	var leftList, rightList []int

	for fileScanner.Scan() {
		x := fileScanner.Text()
		y := strings.Split(x, " ")

		leftList = append(leftList, utils.SafeInt(y[0]))
		rightList = append(rightList, utils.SafeInt(y[len(y)-1]))
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	similarityScore := 0

	for _, x := range leftList {
		ocurrencies := 0
		for _, y := range rightList {
			if x == y {
				ocurrencies++
			}
		}
		similarityScore += x * ocurrencies
	}

	fmt.Println("day 1 part two", similarityScore)
}
