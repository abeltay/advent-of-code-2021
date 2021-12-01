package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var count int
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	prev := nextNum(scanner.Text())
	for scanner.Scan() {
		s := scanner.Text()
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		if prev < i {
			count++
		}
		prev = i
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Count is: %d\n", count)
}

func nextNum(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
