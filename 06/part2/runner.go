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
	var days [9]int
	for i := range population {
		days[population[i]]++
	}
	for i := 0; i < 256; i++ {
		temp := days[0]
		for j := 1; j < len(days); j++ {
			days[j-1] = days[j]
		}
		days[8] = temp
		days[6] += temp
		// fmt.Printf("After %d day: %v\n", i+1, days)
	}
	var total int
	for _, v := range days {
		total += v
	}
	return total
}
