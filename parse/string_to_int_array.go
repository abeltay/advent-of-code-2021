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
