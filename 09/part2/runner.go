package aoc

import (
	"log"
	"sort"

	"github.com/abeltay/advent-of-code-2021/parse"
)

// Runner runs the algorithm to get the answer
func Runner(data []string) int {
	arr := make([][]int, len(data))
	for i := range data {
		iarr, err := parse.ContinuousStringToIntArray(data[i])
		if err != nil {
			log.Fatal(err)
		}
		arr[i] = iarr
	}
	/*
		for i := range arr {
			fmt.Println(arr[i])
		}
	*/
	var sizes []int
	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] >= 9 {
				continue
			}
			sizes = append(sizes, crawl(arr, i, j))
		}
	}
	sort.Ints(sizes)
	// fmt.Println(sizes)
	count := 1
	for i := 0; i < 3; i++ {
		count = count * sizes[len(sizes)-1-i]
	}
	return count
}

func crawl(arr [][]int, x, y int) int {
	if x < 0 || y < 0 || x >= len(arr) || y >= len(arr[x]) || arr[x][y] >= 9 {
		return 0
	}
	arr[x][y] = 9
	count := 1
	count += crawl(arr, x-1, y)
	count += crawl(arr, x+1, y)
	count += crawl(arr, x, y-1)
	count += crawl(arr, x, y+1)
	return count
}
