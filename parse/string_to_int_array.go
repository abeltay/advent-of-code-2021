package parse

import (
	"strconv"
	"strings"
)

// StringToIntArray converts comma-separated values "3,4,3,1,2" into int array []int{3,4,3,1,2}
func StringToIntArray(input string) ([]int, error) {
	sarr := strings.Split(input, ",")
	arr := make([]int, len(sarr))
	for i := range sarr {
		num, err := strconv.Atoi(sarr[i])
		if err != nil {
			return nil, err
		}
		arr[i] = num
	}
	return arr, nil
}

// ContinuousStringToIntArray converts a string "2199943210" into int array []int{2,1,9,9,9,4,3,2,1,0}
func ContinuousStringToIntArray(input string) ([]int, error) {
	arr := make([]int, len(input))
	for i := range input {
		num, err := strconv.Atoi(string(input[i]))
		if err != nil {
			return nil, err
		}
		arr[i] = num
	}
	return arr, nil
}
