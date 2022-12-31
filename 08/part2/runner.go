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
		s := strings.Split(out[0], " ")
		decoded := decode(s)
		// fmt.Printf("%v\n", decoded)
		s = strings.Split(out[1], " ")
		multiplier := 1000
		var num int
		for _, v := range s {
			// fmt.Printf("%v\n", v)
			for j := range decoded {
				if similarMatch(v, decoded[j]) {
					// fmt.Printf("%s %s\n", v, decoded[j])
					num += j * multiplier
					multiplier = multiplier / 10
				}
			}
		}
		// fmt.Printf("%v\n", num)
		count += num
	}
	return count
}

func decode(data []string) []string {
	var d [10]string
	for i := range data {
		switch len(data[i]) {
		case light1:
			d[1] = data[i]
		case light4:
			d[4] = data[i]
		case light7:
			d[7] = data[i]
		case light8:
			d[8] = data[i]
		}
	}
	for i := range data {
		// 2, 3 or 5
		if len(data[i]) == 5 {
			if containsAll(data[i], d[7]) {
				d[3] = data[i]
				continue
			}
			if num := countNotExist(data[i], d[4]); num == 1 {
				d[5] = data[i]
			} else {
				d[2] = data[i]
			}
		}
		// 0, 6 or 9
		if len(data[i]) == 6 {
			if containsAll(data[i], d[4]) {
				d[9] = data[i]
				continue
			}
			if containsAll(data[i], d[1]) {
				d[0] = data[i]
			} else {
				d[6] = data[i]
			}
		}
	}
	return d[:]
}

func containsAll(s, substr string) bool {
	for i := range substr {
		if !strings.ContainsRune(s, rune(substr[i])) {
			return false
		}
	}
	return true
}

func countNotExist(s, substr string) int {
	var diff int
	for i := range substr {
		if !strings.ContainsRune(s, rune(substr[i])) {
			diff++
		}
	}
	return diff
}

func similarMatch(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if !strings.ContainsRune(s2, rune(s1[i])) {
			return false
		}
	}
	return true
}
