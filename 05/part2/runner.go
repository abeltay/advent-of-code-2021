package part2

import (
	"log"
	"strings"

	"github.com/abeltay/advent-of-code-2021/parse"
)

type coord struct {
	x, y int
}

// Runner runs the algorithm to get the answer
func Runner(data []string) int {
	var (
		start, end []coord
		maxX, maxY int
	)
	for i := range data {
		sarr := strings.Split(data[i], " -> ")

		startCoord, err := parse.StringToIntArray(sarr[0])
		if err != nil {
			log.Fatal(err)
		}
		start = append(start, coord{startCoord[0], startCoord[1]})
		if startCoord[0] > maxX {
			maxX = startCoord[0]
		}
		if startCoord[1] > maxY {
			maxY = startCoord[1]
		}

		endCoord, err := parse.StringToIntArray(sarr[1])
		if err != nil {
			log.Fatal(err)
		}
		end = append(end, coord{endCoord[0], endCoord[1]})
		if endCoord[0] > maxX {
			maxX = endCoord[0]
		}
		if endCoord[1] > maxY {
			maxY = endCoord[1]
		}
	}
	// fmt.Println(start)
	// fmt.Println(end)

	maxX++
	maxY++
	ocean := make([][]int, maxY)
	for i := range ocean {
		ocean[i] = make([]int, maxX)
	}

	for i := range start {
		if start[i].x == end[i].x {
			j := start[i].y
			k := end[i].y
			if start[i].y > end[i].y {
				j = end[i].y
				k = start[i].y
			}
			for j <= k {
				ocean[j][start[i].x]++
				j++
			}
			continue
		}
		if start[i].y == end[i].y {
			j := start[i].x
			k := end[i].x
			if start[i].x > end[i].x {
				j = end[i].x
				k = start[i].x
			}
			for j <= k {
				ocean[start[i].y][j]++
				j++
			}
			continue
		}
		drawDiagonal(ocean, start[i], end[i])
	}
	/*
		for i := range arr {
			fmt.Println(arr[i])
		}
	*/
	var count int
	for i := range ocean {
		for j := range ocean[i] {
			if ocean[i][j] > 1 {
				count++
			}
		}
	}
	return count
}

func drawDiagonal(ocean [][]int, start, end coord) {
	if start.y > end.y {
		drawDiagonal(ocean, end, start)
		return
	}
	for i, j := start.y, start.x; i <= end.y; i++ {
		// fmt.Printf("coords: %v %v %d %d\n", start, end, i, j)
		ocean[i][j]++
		if start.x < end.x {
			j++
		} else {
			j--
		}
	}
}
