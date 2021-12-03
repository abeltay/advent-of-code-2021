package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	count      [12]int
	totalCount int
}

// Process processes the input
func (f *Runner) Process(input string) {
	f.totalCount++
	for i := 0; i < len(input); i++ {
		if input[i] == '1' {
			f.count[i]++
		}
	}
}

// Answer prints the answer
func (f Runner) Answer() int64 {
	var gamma, epsilon string
	half := f.totalCount / 2
	for i := 0; i < len(f.count); i++ {
		if f.count[i] > half {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	g, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	e, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return g * e
}
