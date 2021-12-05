package output

import "testing"

func TestTwoDimensionalIntArray(t *testing.T) {
	arr := [][]int{
		{0, 1},
		{1, 1},
	}
	TwoDimensionalIntArray(arr)
}
