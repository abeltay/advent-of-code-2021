package aoc

import (
	"container/list"
	"math"
	"strings"
)

// Runner runs the algorithm to get the answer
func Runner(data []string) int {
	var startSeq string
	var parseSeq bool
	pairs := make(map[string]string)
	for i := range data {
		if data[i] == "" {
			parseSeq = true
			continue
		}
		if !parseSeq {
			startSeq = data[i]
			continue
		}
		s := strings.Split(data[i], " -> ")
		pairs[s[0]] = s[1]
	}
	// ele := poly.Front()
	// for ele != nil {
	// 	s := ele.Value.(string)
	// 	fmt.Println(s)
	// 	ele = ele.Next()
	// }
	// fmt.Println(pairs)
	pairsM := make(map[string]int)
	counter := make(map[string]int)
	counter[string(startSeq[0])]++
	for i := 1; i < len(startSeq); i++ {
		counter[string(startSeq[i])]++
		p := string(startSeq[i-1]) + string(startSeq[i])
		pairsM[p]++
	}
	for n := 0; n < 40; n++ {
		newPairsM := make(map[string]int)
		for k := range pairsM {
			toAdd := pairs[k]
			if toAdd != "" {
				newPairsM[string(k[0])+toAdd] += pairsM[k]
				newPairsM[toAdd+string(k[1])] += pairsM[k]
				counter[toAdd] += pairsM[k]
			}
		}
		pairsM = newPairsM
	}
	// fmt.Println(pairsM)
	// fmt.Println(counter)
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

func morph(pairs [][]string, i1, i2 string, counter map[string]int) {
	poly := list.New()
	poly.PushBack(i1)
	poly.PushBack(i2)
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
	}
	ele := poly.Front()
	for ele != nil {
		s := ele.Value.(string)
		counter[s]++
		ele = ele.Next()
	}
}
