package utils

import (
	"strconv"
	"strings"
)

func SafeInt(s string) int {
	i, _ := strconv.Atoi(strings.TrimPrefix(s, "0"))
	return i
}
