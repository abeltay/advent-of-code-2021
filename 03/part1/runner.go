package part1

import (
	"log"
	"strconv"
)

// Runner runs the algorithm to get the answer
func Runner(data []string) int {
	var totalCount int
	count := make([]int, len(data[0]))
	for i := range data {
		totalCount++
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] == '1' {
				count[j]++
			}
		}
	}
	var gamma, epsilon string
	half := totalCount / 2
	for i := 0; i < len(count); i++ {
		if count[i] > half {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	g, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	e, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(g * e)
}
