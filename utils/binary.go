package utils

import (
	"fmt"
	"strconv"
)

func TrimmedBinary(length int, step int) string {
	/*
		This method is good for getting all the combinations in a
		fixed length if there are only two states.

		length: 3, step: 0	=>	000
		length: 3, step: 1	=>	001
		length: 3, step: 2	=>	010
		length: 3, step: 3	=>	011

		length: 5, step: 3	=>	00011
	*/
	binaryStr := strconv.FormatInt(int64(step), 2)
	if len(binaryStr) > length {
		binaryStr = binaryStr[len(binaryStr)-length:]
	}

	// Pad with zeros if necessary
	return fmt.Sprintf("%0"+strconv.FormatInt(int64(length), 10)+"s", binaryStr)
}
