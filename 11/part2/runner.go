package aoc

import (
	"container/list"
	"log"

	"github.com/abeltay/advent-of-code-2021/parse"
)

type coord struct {
	x, y int
}

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
	var step int
	for !isDone(arr) {
		queue := list.New()
		for j := range arr {
			for k := range arr[j] {
				arr[j][k]++
				if arr[j][k] > 9 {
					queue.PushBack(coord{x: j, y: k})
					arr[j][k] = 0
				}
			}
		}
		ele := queue.Front()
		for ele != nil {
			queue.Remove(ele)
			v := ele.Value.(coord)
			flash(arr, v.x, v.y, queue)
			ele = queue.Front()
		}
		step++
	}
	return step
}

func flash(arr [][]int, x, y int, queue *list.List) {
	// fmt.Printf("coords: %d %d\n", x, y)
	for i := x - 1; i <= x+1; i++ {
		if i < 0 || i >= len(arr) {
			continue
		}
		for j := y - 1; j <= y+1; j++ {
			if j < 0 || j >= len(arr[i]) {
				continue
			}
			if arr[i][j] == 0 {
				continue
			}
			arr[i][j]++
			if arr[i][j] > 9 {
				queue.PushBack(coord{x: i, y: j})
				arr[i][j] = 0
			}
		}
	}
}

func isDone(arr [][]int) bool {
	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] != 0 {
				return false
			}
		}
	}
	return true
}
