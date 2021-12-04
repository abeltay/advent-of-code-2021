package part1

import (
	"log"
	"strconv"
	"strings"
)

// Runner runs the algorithm to get the answer
func Runner(data []string) int {
	sarr := strings.Split(data[0], ",")
	var input []int
	for _, v := range sarr {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, i)
	}
	// fmt.Printf("%v\n", input)

	// Construct arrays
	var (
		boards [][][]int
		checks [][][]bool
		row    int
		bnum   = -1
	)

	for i := 1; i < len(data); i++ {
		if data[i] == "" {
			var board [][]int
			boards = append(boards, board)
			var check [][]bool
			checks = append(checks, check)
			bnum++
			row = 0
			continue
		}
		arr := strings.Split(data[i], " ")
		var iarr []int
		for _, v := range arr {
			i, err := strconv.Atoi(v)
			if err != nil {
				continue
			}
			iarr = append(iarr, i)
		}
		boards[bnum] = append(boards[bnum], iarr)
		check := make([]bool, 5)
		checks[bnum] = append(checks[bnum], check)
		row++
	}
	// fmt.Printf("%v\n", boards)
	// fmt.Printf("%v\n", checks)

	boardsWon := make([]bool, len(boards))
	for ki := range input {
		// fmt.Println(input[ki])
		last := addInput(boards, checks, input[ki], boardsWon)
		var completedCount int
		for _, complete := range boardsWon {
			if complete {
				completedCount++
			}
		}
		if completedCount == len(boards) {
			return calculateWin(input[ki], boards[last], checks[last])
		}
	}
	return 0
}

func addInput(boards [][][]int, checks [][][]bool, input int, boardsWon []bool) int {
	var lastCompleted int
	for bnum := range boards {
		// Mark location
		if boardsWon[bnum] {
			continue
		}
		for x := range boards[bnum] {
			for y := range boards[bnum][x] {
				if checks[bnum][x][y] == true {
					continue
				}
				if boards[bnum][x][y] == input {
					checks[bnum][x][y] = true
					// Check win condition
					if horizontalWin(checks[bnum][x]) || verticalWin(checks[bnum], y) {
						boardsWon[bnum] = true
						lastCompleted = bnum
					}
				}
			}
		}
	}
	return lastCompleted
}

func horizontalWin(row []bool) bool {
	for _, v := range row {
		if !v {
			return false
		}
	}
	return true
}

func verticalWin(board [][]bool, col int) bool {
	for _, v := range board {
		if !v[col] {
			return false
		}
	}
	return true
}

func calculateWin(input int, board [][]int, check [][]bool) int {
	// fmt.Printf("input: %d\nboard: %v\ncheck: %v\n", input, board, check)
	var total int
	for x := range check {
		for y := range check[x] {
			if !check[x][y] {
				// fmt.Printf("adding %d\n", board[x][y])
				total += board[x][y]
			}
		}
	}
	return total * input
}
