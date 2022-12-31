package aoc

import (
	"container/heap"
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
	arr = mapBuilder(arr)
	// for i := range arr {
	// 	fmt.Println(arr[i])
	// }
	riskArr := make([][]int, len(arr))
	for i := range riskArr {
		riskArr[i] = make([]int, len(arr[i]))
		for j := range riskArr[i] {
			riskArr[i][j] = math.MaxInt64
		}
	}
	riskArr[0][0] = 0
	riskArr[0][1] = arr[0][1]
	riskArr[1][0] = arr[1][0]
	h := &nodeHeap{
		node{
			x:    0,
			y:    1,
			risk: riskArr[0][1],
		},
		node{
			x:    1,
			y:    0,
			risk: riskArr[1][0],
		},
	}
	heap.Init(h)
	for h.Len() > 0 {
		calc(arr, riskArr, h)
	}
	return riskArr[len(riskArr)-1][len(riskArr[0])-1]
}
func calc(arr [][]int, riskArr [][]int, h *nodeHeap) {
	n := heap.Pop(h).(node)
	x := n.x
	y := n.y
	if x == len(arr)-1 && y == len(arr[0])-1 {
		return
	}
	pushHeap(h, arr, riskArr, riskArr[x][y], x+1, y)
	pushHeap(h, arr, riskArr, riskArr[x][y], x-1, y)
	pushHeap(h, arr, riskArr, riskArr[x][y], x, y+1)
	pushHeap(h, arr, riskArr, riskArr[x][y], x, y-1)
}

func pushHeap(n *nodeHeap, arr [][]int, riskArr [][]int, risk, x, y int) {
	if x < 0 || x >= len(arr) || y < 0 || y >= len(arr[0]) {
		return
	}
	newRisk := risk + arr[x][y]
	if riskArr[x][y] <= newRisk {
		return
	}
	riskArr[x][y] = newRisk
	if riskArr[x][y] >= riskArr[len(riskArr)-1][len(riskArr[0])-1] {
		return
	}
	heap.Push(n, node{x: x, y: y, risk: riskArr[x][y]})
}

func mapBuilder(arr [][]int) [][]int {
	newArr := make([][]int, len(arr)*5)
	for i := range newArr {
		newArr[i] = make([]int, len(arr[0])*5)
		for j := range newArr[i] {
			factor := j/len(arr[0]) + i/len(arr)
			oldI := i % len(arr)
			oldJ := j % len(arr[0])
			newArr[i][j] = (arr[oldI][oldJ] + factor)
			for newArr[i][j] > 9 {
				newArr[i][j] -= 9
			}
		}
	}
	return newArr
}

type node struct {
	x, y, risk int
}

// An nodeHeap is a min-heap of nodes.
type nodeHeap []node

func (h nodeHeap) Len() int           { return len(h) }
func (h nodeHeap) Less(i, j int) bool { return h[i].risk < h[j].risk }
func (h nodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *nodeHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(node))
}

func (h *nodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
