package part1

import (
	"strings"
)

const (
	light1 = 2
	light4 = 4
	light7 = 3
	light8 = 7
)

// Runner runs the algorithm to get the answer
func Runner(data []string) int {
	var count int
	for i := range data {
		out := strings.Split(data[i], " | ")
		s := strings.Split(out[1], " ")
		for j := range s {
			switch len(s[j]) {
			case light1:
				fallthrough
			case light4:
				fallthrough
			case light7:
				fallthrough
			case light8:
				count++
			}
		}
		// fmt.Printf("%+v\n", s)
	}
	return count
}
