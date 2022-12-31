package aoc

import (
	"log"
	"strconv"
	"strings"
)

// Runner runs the algorithm to get the answer
func Runner(data []string) int {
	s := strings.Split(data[0], " ")
	targetX, err := getRange(s[2])
	if err != nil {
		log.Fatal(err)
	}
	targetY, err := getRange(s[3])
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(targetX)
	// fmt.Println(targetY)
	var x, y, ans int
	for x = 1; x <= targetX[1]; x++ {
		for y = targetY[0]; y < 10000; y++ {
			// fmt.Printf("x: %d, y:%d\n", x, y)
			if calcHeight(targetX, targetY, x, y) {
				// fmt.Printf("x: %d, y:%d\n", x, y)
				ans++
			}
		}
	}
	return ans
}

func getRange(s string) ([]int, error) {
	arr := strings.Split(s, "=")
	arr = strings.Split(arr[1], "..")
	start, err := strconv.Atoi(arr[0])
	if err != nil {
		return nil, err
	}
	arr[1] = strings.TrimRight(arr[1], ",")
	end, err := strconv.Atoi(arr[1])
	if err != nil {
		return nil, err
	}
	return []int{start, end}, nil
}

func calcHeight(targetX, targetY []int, x, y int) bool {
	veloX := x
	veloY := y

	for y >= targetY[0] {
		if x >= targetX[0] && x <= targetX[1] && y >= targetY[0] && y <= targetY[1] {
			return true
		}

		// fmt.Printf("x: %d, y:%d\n", x, y)
		if veloX > 0 {
			veloX--
		}
		veloY--

		x += veloX
		y += veloY
		if x > targetX[1] {
			break
		}
		if y < targetY[0] {
			break
		}
	}
	return false
}
