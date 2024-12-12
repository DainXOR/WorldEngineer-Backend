package types

type Predicate[T any] func(T) bool

func Filter[T any](slice []T, predicate Predicate[T]) []T {
	var result []T
	for _, value := range slice {
		if predicate(value) {
			result = append(result, value)
		}
	}
	return result
}
func Map[T, U any](slice []T, mapper func(T) U) []U {
	var result []U
	for _, value := range slice {
		result = append(result, mapper(value))
	}
	return result
}
func Reduce[T, U any](slice []T, reducer func(U, T) U, initial U) U {
	result := initial
	for _, value := range slice {
		result = reducer(result, value)
	}
	return result
}
