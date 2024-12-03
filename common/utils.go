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
