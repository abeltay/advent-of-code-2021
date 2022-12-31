package aoc

import (
	"sort"
)

// Runner runs the algorithm to get the answer
func Runner(data []string) int {
	var counts []int
	for i := range data {
		stack := make([]byte, 0, len(data[i]))
		var corrupted bool
		for j := range data[i] {
			switch data[i][j] {
			case '(':
				fallthrough
			case '[':
				fallthrough
			case '{':
				fallthrough
			case '<':
				stack = append(stack, data[i][j])
			case ')':
				n := len(stack) - 1 // Top element
				if stack[n] != '(' {
					corrupted = true
					break
				}
				stack = stack[:n]
			case ']':
				n := len(stack) - 1 // Top element
				if stack[n] != '[' {
					corrupted = true
					break
				}
				stack = stack[:n]
			case '}':
				n := len(stack) - 1 // Top element
				if stack[n] != '{' {
					corrupted = true
					break
				}
				stack = stack[:n]
			case '>':
				n := len(stack) - 1 // Top element
				if stack[n] != '<' {
					corrupted = true
					break
				}
				stack = stack[:n]
			}
			if corrupted {
				break
			}
		}
		if !corrupted {
			// Calculate score
			var count int
			for i := len(stack) - 1; i >= 0; i-- {
				count *= 5
				switch stack[i] {
				case '(':
					count++
				case '[':
					count += 2
				case '{':
					count += 3
				case '<':
					count += 4
				}
			}
			counts = append(counts, count)
		}
	}
	sort.Ints(counts)
	// fmt.Println(counts)
	return counts[len(counts)/2]
}
