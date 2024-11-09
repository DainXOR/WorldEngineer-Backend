package utils

func Index[S ~[]E, E comparable](slice S, element E) int {
	for i, value := range slice {
		if element == value {
			return i
		}
	}
	return -1
}
func Contains[S ~[]E, E comparable](slice S, element E) bool {
	return Index(slice, element) >= 0
}

func Map[S ~[]E, E, R any](slice S, mutator func(E) R) []R {
	result := make([]R, len(slice))
	for i, value := range slice {
		result[i] = mutator(value)
	}
	return result
}
func MMap[K comparable, M ~map[K]V, V, R any](dict M, mutator func(K, V) R) map[K]R {
	result := make(map[K]R, len(dict))
	for key, value := range dict {
		result[key] = mutator(key, value)
	}
	return result
}

func Filter[S ~[]E, E any](slice S, predicate func(E) bool) []E {
	result := make([]E, 0, len(slice))
	for _, value := range slice {
		if predicate(value) {
			result = append(result, value)
		}
	}
	return result
}

func Reduce[S ~[]E, E, R any](slice S, reducer func(R, E) R, initial R) R {
	result := initial
	for _, value := range slice {
		result = reducer(result, value)
	}
	return result
}
