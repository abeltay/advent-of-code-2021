package part2

import (
	"fmt"
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
	size := make([]int, len(sarr))
	for i := range size {
		size[i] = 1
	}
	for i := 0; i < 256; i++ {
		var baby int
		for j := range population {
			if population[j] <= 0 {
				baby += size[j]
				population[j] = 6
				continue
			}
			population[j]--
		}
		if baby == 0 {
			continue
		}
		population = append(population, 8)
		size = append(size, baby)
		// fmt.Printf("After %d day: \n", i+1)
		fmt.Printf("After %d day: %v\n", i+1, size)
	}
	var total int
	for i := range size {
		total += size[i]
	}
	return total
}
