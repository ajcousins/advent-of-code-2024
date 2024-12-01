package utils

func GetSmallestValue(vals []int) int {
	smallest := vals[0]
	for _, v := range vals {
		if v < smallest {
			smallest = v
		}
	}

	return smallest
}
