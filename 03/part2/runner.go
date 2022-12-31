package part2

import (
	"log"
	"strconv"
	"strings"
)

// Runner runs the algorithm to get the answer
func Runner(data []string) int {
	var prefix, val string
	for j := 0; j < len(data[0]); j++ {
		var count, total int
		for i := 0; i < len(data); i++ {
			if strings.HasPrefix(data[i], prefix) {
				total++
				if data[i][j] == '1' {
					count++
				}
				val = data[i]
			}
		}
		if total == 1 {
			break
		}
		half := float64(total) / 2
		if float64(count) >= half {
			prefix += "1"
		} else {
			prefix += "0"
		}
	}
	oxygen, err := strconv.ParseInt(val, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(oxygen, "is the oxygen generator rating")

	prefix, val = "", ""
	for j := 0; j < len(data[0]); j++ {
		var count, total int
		for i := 0; i < len(data); i++ {
			if strings.HasPrefix(data[i], prefix) {
				total++
				if data[i][j] == '1' {
					count++
				}
				val = data[i]
			}
		}
		if total == 1 {
			break
		}
		half := float64(total) / 2
		if float64(count) >= half {
			prefix += "0"
		} else {
			prefix += "1"
		}
	}
	co2, err := strconv.ParseInt(val, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(co2, "is the CO2 scrubber rating")

	return int(oxygen * co2)
}
