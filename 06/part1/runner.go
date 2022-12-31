package part1

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
