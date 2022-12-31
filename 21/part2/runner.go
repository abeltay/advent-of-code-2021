package aoc

import (
	"log"
	"strconv"
	"strings"
)

// Runner runs the algorithm to get the answer
func Runner(data []string) int {
	s := strings.Split(data[0], ": ")
	p1, err := strconv.Atoi(s[1])
	if err != nil {
		log.Fatal(err)
	}
	s = strings.Split(data[1], ": ")
	p2, err := strconv.Atoi(s[1])
	if err != nil {
		log.Fatal(err)
	}
	u := universe{
		p1pos: p1,
		p2pos: p2,
	}
	one, two := roller(u)
	// fmt.Println(one, two)
	ans := one
	if two > ans {
		ans = two
	}
	return ans
}

func normalise(position int) int {
	return ((position - 1) % 10) + 1
}

type universe struct {
	p1pos   int
	p2pos   int
	p1score int
	p2score int
}

var cache = make(map[universe][]int)

func roller(orig universe) (int, int) {
	if orig.p2score >= 21 {
		// fmt.Println(orig.p1score, orig.p2score)
		return 0, 1
	}
	arr, ok := cache[orig]
	if ok {
		// fmt.Println("found", orig, arr)
		one, two := arr[0], arr[1]
		return one, two
	}
	var p1, p2 int
	for i1 := 1; i1 <= 3; i1++ {
		for i2 := 1; i2 <= 3; i2++ {
			for i3 := 1; i3 <= 3; i3++ {
				newPos := normalise(orig.p1pos + i1 + i2 + i3)
				newScore := orig.p1score + newPos
				newU := universe{
					p1pos:   orig.p2pos,
					p2pos:   newPos,
					p1score: orig.p2score,
					p2score: newScore,
				}
				one, two := roller(newU)
				p1 += two
				p2 += one
			}
		}
	}
	// fmt.Println(p1, p2)
	cache[orig] = []int{p1, p2}
	return p1, p2
}
