package aoc

import (
	"log"
	"strconv"
)

// Runner runs the algorithm to get the answer
func Runner(data []string) int {
	var arr [][]int
	for i := range data {
		var iarr []int
		for j := range data[i] {
			num, err := strconv.Atoi(string(data[i][j]))
			if err != nil {
				log.Fatal(err)
			}
			iarr = append(iarr, num)
		}
		arr = append(arr, iarr)
	}
	/*
		for i := range arr {
			fmt.Println(arr[i])
		}
	*/
	var count int
	for i := range arr {
		for j := range arr[i] {
			if lowestCardinal(arr, i, j) {
				// fmt.Println(arr[i][j])
				count += arr[i][j] + 1
			}
		}
	}
	return count
}

func lowestCardinal(arr [][]int, x, y int) bool {
	if x-1 >= 0 && arr[x-1][y] <= arr[x][y] {
		return false
	}
	if x+1 < len(arr) && arr[x+1][y] <= arr[x][y] {
		return false
	}
	if y-1 >= 0 && arr[x][y-1] <= arr[x][y] {
		return false
	}
	if y+1 < len(arr[x]) && arr[x][y+1] <= arr[x][y] {
		return false
	}
	return true
}
