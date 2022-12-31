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
	file, err := os.Open("02/input.txt")
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

	fmt.Printf("Answer is: %d\n", r.Answer())
}

// Runner holds the variables and methods for a run
type Runner struct {
	right int
	depth int
}

// Process processes the input
func (f *Runner) Process(input string) {
	s := strings.Split(input, " ")
	i, err := strconv.Atoi(s[1])
	if err != nil {
		log.Fatal(err)
	}
	switch s[0] {
	case "forward":
		f.right += i
	case "down":
		f.depth += i
	case "up":
		f.depth -= i
	}
}

// Answer prints the answer
func (f Runner) Answer() int {
	return f.right * f.depth
}
