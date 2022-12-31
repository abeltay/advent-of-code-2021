package aoc

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Runner runs the algorithm to get the answer
func Runner(data []string) {
	// Work backwards through the data to find answer
	fmt.Println(valid(data, 99911993949684))
	fmt.Println(valid(data, 62911941716111))
	// for start := 99999999999999; start >= 11111111111111; start-- {
	// 	// fmt.Println(start)
	// 	if valid(data, start) {
	// 		return start
	// 	}
	// }
}

func valid(data []string, start int) bool {
	// Check if start has zeroes
	str := fmt.Sprintf("%014d", start)
	for j := range str {
		if str[j] == '0' {
			return false
		}
	}
	var pointer int
	var w, x, y, z int
	for i := range data {
		instruc := strings.Split(data[i], " ")
		if instruc[0] == "inp" {
			num, err := strconv.Atoi(string(str[pointer]))
			if err != nil {
				log.Fatal(err)
			}
			pointer++
			w = num
			continue
		}
		process(instruc, &w, &x, &y, &z)
		// fmt.Println(instruc, w, x, y, z)
	}
	if z == 0 {
		return true
	}
	// fmt.Println(w, x, y, z)
	return false
}

func process(ins []string, w, x, y, z *int) {
	num, err := strconv.Atoi(ins[2])
	if err != nil {
		// Means it's a character
		b := charToBuffer(ins[2], w, x, y, z)
		num = *b
	}
	a := charToBuffer(ins[1], w, x, y, z)
	switch ins[0] {
	case "add":
		*a += num
	case "mul":
		*a *= num
	case "div":
		if num != 0 {
			*a = *a / num
		}
	case "mod":
		if *a >= 0 && num > 0 {
			*a = *a % num
		}
	case "eql":
		if *a == num {
			*a = 1
		} else {
			*a = 0
		}
	default:
		log.Fatal("invalid ins", ins[0])
	}
}

func charToBuffer(char string, w, x, y, z *int) *int {
	switch char {
	case "w":
		return w
	case "x":
		return x
	case "y":
		return y
	case "z":
		return z
	default:
		return nil
	}
}
