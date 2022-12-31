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

	var cubes []cube
	for k := 0; k < len(arr); k++ {
		var intersects []cube
		// fmt.Println("-----", k)
		for i := range cubes {
			// Nullify all intersecting cubes
			c := createIntersect(cubes[i], arr[k])
			if validCube(c) {
				c.on = !cubes[i].on
				// fmt.Println("nullifying", cubes[i], arr[k], c)
				intersects = append(intersects, c)
			}
		}
		cubes = append(cubes, intersects...)
		if arr[k].on {
			cubes = append(cubes, arr[k])
		}
		// for j := range cubes {
		// 	fmt.Println(cubes[j])
		// }
	}
	var ans int
	for i := range cubes {
		// fmt.Println(cubes[i])
		volume := (cubes[i].toX - cubes[i].fromX + 1) * (cubes[i].toY - cubes[i].fromY + 1) * (cubes[i].toZ - cubes[i].fromZ + 1)
		if cubes[i].on {
			ans += volume
		} else {
			ans -= volume
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

func createIntersect(c1, c2 cube) cube {
	x1 := c1.fromX
	if c1.fromX < c2.fromX {
		x1 = c2.fromX
	}
	x2 := c1.toX
	if c1.toX > c2.toX {
		x2 = c2.toX
	}
	y1 := c1.fromY
	if c1.fromY < c2.fromY {
		y1 = c2.fromY
	}
	y2 := c1.toY
	if c1.toY > c2.toY {
		y2 = c2.toY
	}
	z1 := c1.fromZ
	if c1.fromZ < c2.fromZ {
		z1 = c2.fromZ
	}
	z2 := c1.toZ
	if c1.toZ > c2.toZ {
		z2 = c2.toZ
	}
	out := cube{
		fromX: x1,
		toX:   x2,
		fromY: y1,
		toY:   y2,
		fromZ: z1,
		toZ:   z2,
		on:    false,
	}
	// fmt.Println(c1, c2, out)
	return out
}

func validCube(c cube) bool {
	return c.fromX <= c.toX && c.fromY <= c.toY && c.fromZ <= c.toZ
}
