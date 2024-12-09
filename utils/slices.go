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

func MapKeysToSlice[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}

func DeduplicateSlice[T comparable](slice []T) []T {
	set := map[T]bool{}
	for _, el := range slice {
		set[el] = true
	}

	return MapKeysToSlice(set)
}

func ReverseSlice[T any](slice []T) []T {
	reversed := make([]T, len(slice))
	for i, v := range slice {
		reversed[(len(slice)-1)-i] = v
	}
	return reversed
}

func FilterElement[T comparable](slice []T, el T) []T {
	newSlice := []T{}
	for _, sliceEl := range slice {
		if el == sliceEl {
			continue
		}
		newSlice = append(newSlice, sliceEl)
	}
	return newSlice
}
