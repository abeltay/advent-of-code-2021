package part2

import (
	"log"
	"strconv"
	"strings"
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

		startCoord := strings.Split(sarr[0], ",")
		x, err := strconv.Atoi(startCoord[0])
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.Atoi(startCoord[1])
		if err != nil {
			log.Fatal(err)
		}
		start = append(start, coord{x, y})
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}

		endCoord := strings.Split(sarr[1], ",")
		x, err = strconv.Atoi(endCoord[0])
		if err != nil {
			log.Fatal(err)
		}
		y, err = strconv.Atoi(endCoord[1])
		if err != nil {
			log.Fatal(err)
		}
		end = append(end, coord{x, y})
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
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
