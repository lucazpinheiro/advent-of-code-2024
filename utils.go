package utils

import (
	"strconv"
	"strings"
)

func SafeInt(s string) int {
	i, _ := strconv.Atoi(strings.TrimPrefix(s, "0"))
	return i
}

func StringToInt(arr []string) []int {
	intArr := []int{}

	for _, v := range arr {
		intArr = append(intArr, SafeInt(v))
	}

	return intArr
}
