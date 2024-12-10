package common

import "strconv"

func StringToInt(s string) int {
	n, err := strconv.Atoi(s)
	Check(err)
	return n
}

func ConvertStringsToInts(strings []string) []int {
	ints := make([]int, len(strings))
	for i, s := range strings {
		ints[i] = StringToInt(s)
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

func ConcatenateNumbers(a, b int) int {
	strA := strconv.Itoa(a)
	strB := strconv.Itoa(b)

	concatenatedStr := strA + strB

	concatenatedInt, err := strconv.Atoi(concatenatedStr)
	Check(err)

	return concatenatedInt
}
