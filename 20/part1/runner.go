package aoc

import (
	"log"
	"math"
	"strconv"
)

// Runner runs the algorithm to get the answer
func Runner(data []string) int {
	var imageEnh string
	var notfirst bool
	var y int
	var arr []point
	for i := range data {
		if !notfirst {
			imageEnh = data[i]
			notfirst = true
			continue
		}
		if data[i] == "" {
			continue
		}
		var x int
		for j := range data[i] {
			if data[i][j] == '#' {
				arr = append(arr, point{x: x, y: y})
			}
			x++
		}
		y++
	}
	// fmt.Println(imageEnh)
	var canvasLight bool
	for count := 0; count < 2; count++ {
		minX, maxX, minY, maxY := spread(arr)
		var newArr []point
		for x := minX - 1; x <= maxX+1; x++ {
			for y := minY - 1; y <= maxY+1; y++ {
				var s string
				for j := y - 1; j <= y+1; j++ {
					for i := x - 1; i <= x+1; i++ {
						if light(arr, canvasLight, minX, maxX, minY, maxY, i, j) {
							s += "1"
						} else {
							s += "0"
						}
					}
				}
				pointer, err := strconv.ParseInt(s, 2, 64)
				if err != nil {
					log.Fatal(err)
				}
				// fmt.Println(x, y, s, pointer)
				if imageEnh[pointer] == '#' {
					newArr = append(newArr, point{x: x, y: y})
				}
			}
		}
		if canvasLight {
			if imageEnh[511] == '.' {
				canvasLight = false
			}
		} else {
			if imageEnh[0] == '#' {
				canvasLight = true
			}
		}
		arr = newArr
	}
	// for i := range arr {
	// 	fmt.Println(arr[i].x, arr[i].y)
	// }
	return len(arr)
}

type point struct {
	x, y int
}

func spread(arr []point) (int, int, int, int) {
	var minX, maxX, minY, maxY int
	minX, minY = math.MaxInt32, math.MaxInt32
	maxX, maxY = math.MinInt32, math.MinInt32
	for i := range arr {
		if arr[i].x < minX {
			minX = arr[i].x
		}
		if arr[i].x > maxX {
			maxX = arr[i].x
		}
		if arr[i].y < minY {
			minY = arr[i].y
		}
		if arr[i].y > maxY {
			maxY = arr[i].y
		}
	}
	return minX, maxX, minY, maxY
}
func light(arr []point, light bool, minX, maxX, minY, maxY, x, y int) bool {
	if x < minX || x > maxX {
		return light
	}
	if y < minY || y > maxY {
		return light
	}
	for i := range arr {
		if arr[i].x == x && arr[i].y == y {
			return true
		}
	}
	return false
}

type pointSlice []point

func (f pointSlice) Len() int {
	return len(f)
}

func (f pointSlice) Less(i, j int) bool {
	if f[i].x < f[j].x {
		return true
	}
	if f[i].x == f[j].x && f[i].y < f[j].y {
		return true
	}
	return false
}

func (f pointSlice) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
