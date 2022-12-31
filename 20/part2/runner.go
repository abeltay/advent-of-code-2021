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
	var minX, maxX, minY, maxY int
	minX, minY = math.MaxInt32, math.MaxInt32
	maxX, maxY = math.MinInt32, math.MinInt32
	arr := make(map[[2]int]bool)
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
				arr[[2]int{x, y}] = true
				if x < minX {
					minX = x
				}
				if x > maxX {
					maxX = x
				}
				if y < minY {
					minY = y
				}
				if y > maxY {
					maxY = y
				}
			}
			x++
		}
		y++
	}
	// fmt.Println(imageEnh)
	var canvasLight bool
	for count := 0; count < 50; count++ {
		newArr := make(map[[2]int]bool)
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
				pointer, err := strconv.ParseInt(s, 2, 32)
				if err != nil {
					log.Fatal(err)
				}
				// fmt.Println(x, y, s, pointer)
				if imageEnh[pointer] == '#' {
					newArr[[2]int{x, y}] = true
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
		minX--
		maxX++
		minY--
		maxY++
		arr = newArr
		// fmt.Println(count+1, len(arr))
	}
	// for i := range arr {
	// 	fmt.Println(arr[i].x, arr[i].y)
	// }
	return len(arr)
}

type point struct {
	x, y int
}

func light(arr map[[2]int]bool, light bool, minX, maxX, minY, maxY, x, y int) bool {
	if x < minX || x > maxX {
		return light
	}
	if y < minY || y > maxY {
		return light
	}
	return arr[[2]int{x, y}]
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
