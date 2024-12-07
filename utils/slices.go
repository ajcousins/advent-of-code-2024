package utils

import (
	"golang.org/x/exp/slices"
)

func MapSlice[S, T any](input []S, mapper func(S) T) []T {
	result := make([]T, len(input))
	for i, item := range input {
		result[i] = mapper(item)
	}
	return result
}

func RemoveNthElement[S any](input []S, index int) []S {
	return append(input[:index], input[index+1:]...)
}

func ContainsAny[T comparable](slice1 []T, slice2 ...T) bool {
	for _, item := range slice2 {
		if slices.Contains(slice1, item) {
			return true
		}
	}
	return false
}

func AnyInCommon(group1, group2 []string) bool {
	return ContainsAny(group1, group2...)
}

func SwapValues[T comparable](values []T, index int) []T {
	/*
		arr := []int{0, 1, 2, 3, 4}
		s1 := arr[:]   // [0 1 2 3 4]
		s2 := arr[1:]  // [1 2 3 4]
		s3 := arr[:3]  // [0 1 2]
	*/

	start := values[:index-1]
	swapA := values[index-1]
	swapB := values[index]
	end := values[index+1:]
	newSlice := append(append(append(start, swapB), swapA), end...)

	return newSlice
}
