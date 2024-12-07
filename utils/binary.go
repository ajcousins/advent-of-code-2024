package utils

import (
	"fmt"
	"strconv"
)

func TrimmedBaseNNumber(length int, step int, states int) string {
	/*
		This method is good for getting all the combinations in a
		fixed length if there are an "n" number of states.

		length: 3, step: 0, states: 2	=>	000
		length: 3, step: 1, states: 2	=>	001
		length: 3, step: 2, states: 2	=>	010
		length: 3, step: 3, states: 2	=>	011
		length: 5, step: 3, states: 2	=>	00011
	*/
	baseString := strconv.FormatInt(int64(step), states)
	if len(baseString) > length {
		baseString = baseString[len(baseString)-length:]
	}

	// Pad with zeros if necessary
	for len(baseString) < length {
		baseString = fmt.Sprintf("0%v", baseString)
	}

	return baseString
}

func TrimmedBinary(length int, step int) string {
	return TrimmedBaseNNumber(length, step, 2)
}
