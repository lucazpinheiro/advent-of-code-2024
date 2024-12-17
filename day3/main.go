package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	var fullLine string
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fullLine += line
	}

	var sum = parseMemoryLine(fullLine)
	fmt.Println("day 3 part one:", sum)
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

	var fullLine string
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fullLine += line
	}

	r, _ := regexp.Compile(`do\(\)|mul\((\d{1,3}),(\d{1,3})\)|don't\(\)`)
	sequences := r.FindAllString(fullLine, -1)

	var validSequences []string

	prev := ""
	for _, elem := range sequences {
		if elem == "don't()" {
			prev = elem
			continue
		}
		if elem == "do()" {
			prev = elem
			continue
		}
		if prev == "don't()" {
			continue
		}
		validSequences = append(validSequences, elem)
	}

	sum := 0
	for _, mul := range validSequences {
		sum += retrieveValues(mul)
	}
	fmt.Println("day 3 part one:", sum)
}

func parseMemoryLine(line string) int {
	r, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)

	validSequences := r.FindAllString(line, -1)

	if len(validSequences) < 1 {
		return 0
	}

	acc := 0

	for _, mul := range validSequences {
		acc += retrieveValues(mul)
	}

	return acc
}

func retrieveValues(s string) int {
	mul := strings.Split(s, ",")

	first := utils.SafeInt(mul[0][4:])
	second := utils.SafeInt(mul[1][:len(mul[1])-1])

	return first * second
}
