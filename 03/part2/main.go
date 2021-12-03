package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var r Runner
	for scanner.Scan() {
		r.Process(scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(r.Answer(), "is the answer")
}

// Runner holds the variables and methods for a run
type Runner struct {
	everything []string
}

// Process processes the input
func (f *Runner) Process(input string) {
	f.everything = append(f.everything, input)
}

// Answer prints the answer
func (f Runner) Answer() int64 {
	var prefix, val string
	for j := 0; j < len(f.everything[0]); j++ {
		var count, total int
		for i := 0; i < len(f.everything); i++ {
			if strings.HasPrefix(f.everything[i], prefix) {
				total++
				if f.everything[i][j] == '1' {
					count++
				}
				val = f.everything[i]
			}
		}
		if total == 1 {
			break
		}
		half := float64(total) / 2
		if float64(count) >= half {
			prefix += "1"
		} else {
			prefix += "0"
		}
	}
	oxygen, err := strconv.ParseInt(val, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(oxygen, "is the oxygen generator rating")

	prefix, val = "", ""
	for j := 0; j < len(f.everything[0]); j++ {
		var count, total int
		for i := 0; i < len(f.everything); i++ {
			if strings.HasPrefix(f.everything[i], prefix) {
				total++
				if f.everything[i][j] == '1' {
					count++
				}
				val = f.everything[i]
			}
		}
		if total == 1 {
			break
		}
		half := float64(total) / 2
		if float64(count) >= half {
			prefix += "0"
		} else {
			prefix += "1"
		}
	}
	co2, err := strconv.ParseInt(val, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(co2, "is the CO2 scrubber rating")

	return oxygen * co2
}
