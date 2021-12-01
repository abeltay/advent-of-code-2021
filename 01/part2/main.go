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
	prev1 := nextNum(scanner.Text())
	scanner.Scan()
	prev2 := nextNum(scanner.Text())
	scanner.Scan()
	prev3 := nextNum(scanner.Text())
	for scanner.Scan() {
		cur := nextNum(scanner.Text())
		if prev1 < cur {
			count++
		}
		prev1 = prev2
		prev2 = prev3
		prev3 = cur
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
