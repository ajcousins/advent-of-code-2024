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
