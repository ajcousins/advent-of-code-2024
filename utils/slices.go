package utils

func MapSlice[S, T any](input []S, mapper func(S) T) []T {
	result := make([]T, len(input))
	for i, item := range input {
		result[i] = mapper(item)
	}
	return result
}
