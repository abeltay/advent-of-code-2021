package output

import "fmt"

// TwoDimensionalIntArray prints a nice 2D array
func TwoDimensionalIntArray(arr [][]int) {
	for i := range arr {
		fmt.Println(arr[i])
	}
}
