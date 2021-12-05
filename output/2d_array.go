package output

import "fmt"

func TwoDimensionalIntArray(arr [][]int) {
	for i := range arr {
		fmt.Println(arr[i])
	}
}
