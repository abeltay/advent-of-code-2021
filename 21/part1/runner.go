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
	dice := 1
	var rollcount int
	var p1score, p2score int
	var p2turn bool
	for p1score < 1000 && p2score < 1000 {
		var total int
		for count := 0; count < 3; count++ {
			total += dice
			dice = roll(dice)
			rollcount++
		}
		if !p2turn {
			p1 += total
			p1 = normalise(p1)
			p1score += p1
			p2turn = true
		} else {
			p2 += total
			p2 = normalise(p2)
			p2score += p2
			p2turn = false
		}
	}
	// fmt.Println(p1score, p2score, rollcount)
	min := p1score
	if p2score < min {
		min = p2score
	}
	return min * rollcount
}

func roll(dice int) int {
	n := dice + 1
	if n > 100 {
		return n - 100
	}
	return n
}

func normalise(position int) int {
	return ((position - 1) % 10) + 1
}
