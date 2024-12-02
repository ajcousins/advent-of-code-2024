package utils

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
