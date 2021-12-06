package part2

import (
	"log"

	"github.com/abeltay/advent-of-code-2021/parse"
)

// Runner runs the algorithm to get the answer
func Runner(data []string) int {
	population, err := parse.StringToIntArray(data[0])
	if err != nil {
		log.Fatal(err)
	}
	size := make([]int, len(population))
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
		// fmt.Printf("After %d day: %v\n", i+1, size)
	}
	var total int
	for i := range size {
		total += size[i]
	}
	return total
}
