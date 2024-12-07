package utils

import (
	"fmt"
	"strconv"
)

func StringToInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	return int(i)
}

func BinaryToDecimal(binaryStr string) int {
	decimal, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		return 0
	}
	return int(decimal)
}
