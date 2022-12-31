package aoc

import (
	"log"
	"math"

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
	// for i := range arr {
	// 	fmt.Println(arr[i])
	// }
	visited := make([][]bool, len(arr))
	for i := range visited {
		visited[i] = make([]bool, len(arr[i]))
	}
	visited[0][0] = true
	visited[0][1] = true
	visited[1][0] = true
	for i := 2; i < len(arr)*2; i++ {
		for j := 0; j <= i; j++ {
			calc(arr, visited, j, i-j)
		}
	}
	// for i := range visited {
	// 	fmt.Println(visited[i])
	// }
	// for i := range arr {
	// 	fmt.Println(arr[i])
	// }
	return arr[len(arr)-1][len(arr[0])-1]
}
func calc(arr [][]int, visited [][]bool, x, y int) {
	if x < 0 || x >= len(arr) || y < 0 || y >= len(arr[0]) {
		return
	}
	var directions [4]int
	directions[0] = getRisk(arr, visited, x+1, y)
	directions[1] = getRisk(arr, visited, x-1, y)
	directions[2] = getRisk(arr, visited, x, y+1)
	directions[3] = getRisk(arr, visited, x, y-1)

	min := math.MaxInt64
	// fmt.Println(directions)
	for _, v := range directions {
		if min > v {
			min = v
		}
	}
	if min != math.MaxInt64 {
		arr[x][y] += min
	}
	visited[x][y] = true
}

func getRisk(arr [][]int, visited [][]bool, x, y int) int {
	if x >= 0 && x < len(arr) && y >= 0 && y < len(arr[0]) && visited[x][y] {
		return arr[x][y]
	}
	return math.MaxInt64
}
