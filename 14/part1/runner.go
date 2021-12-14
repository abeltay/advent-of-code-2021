package aoc

import (
	"container/list"
	"math"
	"strings"
)

// Runner runs the algorithm to get the answer
func Runner(data []string) int {
	poly := list.New()
	var parseSeq bool
	var pairs [][]string
	for i := range data {
		if data[i] == "" {
			parseSeq = true
			continue
		}
		if !parseSeq {
			for j := range data[i] {
				poly.PushBack(string(data[i][j]))
			}
			continue
		}
		s := strings.Split(data[i], " -> ")
		pairs = append(pairs, s)
	}
	// ele := poly.Front()
	// for ele != nil {
	// 	s := ele.Value.(string)
	// 	fmt.Println(s)
	// 	ele = ele.Next()
	// }
	for steps := 0; steps < 10; steps++ {
		ele := poly.Front().Next()
		for ele != nil {
			s0 := ele.Prev().Value.(string)
			s1 := ele.Value.(string)
			for i := range pairs {
				if pairs[i][0] == s0+s1 {
					poly.InsertBefore(pairs[i][1], ele)
					break
				}
			}
			ele = ele.Next()
		}
		// ele = poly.Front()
		// for ele != nil {
		// 	s := ele.Value.(string)
		// 	fmt.Print(s)
		// 	ele = ele.Next()
		// }
		// fmt.Println()
	}
	counter := make(map[string]int)
	ele := poly.Front()
	for ele != nil {
		s := ele.Value.(string)
		counter[s]++
		ele = ele.Next()
	}
	var max int
	min := math.MaxInt64
	for _, v := range counter {
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}
	return max - min
}
