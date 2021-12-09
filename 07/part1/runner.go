package part1

import (
	"log"
	"math"
	"sort"

	"github.com/abeltay/advent-of-code-2021/parse"
)

// Runner runs the algorithm to get the answer
func Runner(data []string) int {
	loc, err := parse.StringToIntArray(data[0])
	if err != nil {
		log.Fatal(err)
	}
	sort.Ints(loc)
	// fmt.Println(loc)

	last := math.MaxInt64
	for i := 0; i < loc[len(loc)-1]; i++ {
		d := calculateDist(loc, i)
		if d > last {
			// fmt.Println(i)
			return last
		}
		last = d
	}
	return 0
}

func calculateDist(arr []int, point int) int {
	var total int
	for i := range arr {
		diff := arr[i] - point
		if diff < 0 {
			diff = diff * -1
		}
		total += diff
	}
	return total
}
