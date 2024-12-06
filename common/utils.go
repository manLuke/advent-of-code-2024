package common

import "strconv"

func ConvertStringsToInts(strings []string) []int {
	ints := make([]int, len(strings))
	for i, s := range strings {
		n, err := strconv.Atoi(s)
		Check(err)
		ints[i] = n
	}
	return ints
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func CopyMap[K comparable, V any](original map[K]V) map[K]V {
	newMap := make(map[K]V)
	for key, value := range original {
		newMap[key] = value
	}
	return newMap
}
