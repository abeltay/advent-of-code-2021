package aoc

import (
	"log"
	"strconv"
	"strings"
)

// Runner runs the algorithm to get the answer
func Runner(data []string) int {
	arr := make([]cube, len(data))
	for i := range data {
		s := strings.Split(data[i], " ")
		var c cube
		if s[0] == "on" {
			c.on = true
		}
		s2 := strings.Split(s[1], ",")
		one, two := parseRange(s2[0])
		c.fromX = one
		c.toX = two
		one, two = parseRange(s2[1])
		c.fromY = one
		c.toY = two
		one, two = parseRange(s2[2])
		c.fromZ = one
		c.toZ = two
		arr[i] = c
	}
	// for i := range arr {
	// 	fmt.Println(arr[i])
	// }
	m := make(map[[3]int]bool)
	for _, v := range arr {
		for x := v.fromX; x <= v.toX; x++ {
			if x < -50 || x > 50 {
				continue
			}
			for y := v.fromY; y <= v.toY; y++ {
				if y < -50 || y > 50 {
					continue
				}
				for z := v.fromZ; z <= v.toZ; z++ {
					if z < -50 || z > 50 {
						continue
					}
					m[[3]int{x, y, z}] = v.on
				}
			}
		}
		// fmt.Println("----")
		// for k, v := range m {
		// 	if v {
		// 		fmt.Println(k)
		// 	}
		// }
	}
	var ans int
	for _, v := range m {
		if v {
			ans++
		}
	}
	return ans
}

type cube struct {
	fromX int
	toX   int
	fromY int
	toY   int
	fromZ int
	toZ   int
	on    bool
}

// parseRange parses "x=-20..26"
func parseRange(s string) (int, int) {
	s1 := strings.Split(s, "=")
	s2 := strings.Split(s1[1], "..")
	num, err := strconv.Atoi(s2[0])
	if err != nil {
		log.Fatal(err)
	}
	num2, err := strconv.Atoi(s2[1])
	if err != nil {
		log.Fatal(err)
	}
	return num, num2
}
