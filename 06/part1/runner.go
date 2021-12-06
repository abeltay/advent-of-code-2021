package part1

import (
	"log"
	"strconv"
	"strings"
)

// Runner runs the algorithm to get the answer
func Runner(data []string) int {
	sarr := strings.Split(data[0], ",")
	population := make([]int, 0, len(sarr))
	for i := range sarr {
		num, err := strconv.Atoi(sarr[i])
		if err != nil {
			log.Fatal(err)
		}
		population = append(population, num)
	}
	for i := 0; i < 80; i++ {
		var baby int
		for j := range population {
			if population[j] <= 0 {
				baby++
				population[j] = 6
				continue
			}
			population[j]--
		}
		for j := 0; j < baby; j++ {
			population = append(population, 8)
		}
		// fmt.Printf("After %d day: %v\n", i+1, population)
	}
	return len(population)
}
