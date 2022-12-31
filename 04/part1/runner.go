package part1

import (
	"log"
	"strconv"
	"strings"

	"github.com/abeltay/advent-of-code-2021/parse"
)

// Runner runs the algorithm to get the answer
func Runner(data []string) int {
	input, err := parse.StringToIntArray(data[0])
	if err != nil {
		log.Fatal(err)
	}

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

	for ki := range input {
		// fmt.Println(input[ki])
		res := addInput(boards, checks, input[ki])
		if res != 0 {
			// fmt.Printf("%v\n", checks)
			return res
		}
	}
	return 0
}

func addInput(boards [][][]int, checks [][][]bool, input int) int {
	for bnum := range boards {
		// Mark location
		for x := range boards[bnum] {
			for y := range boards[bnum][x] {
				if checks[bnum][x][y] == true {
					continue
				}
				if boards[bnum][x][y] == input {
					checks[bnum][x][y] = true
					// Check win condition
					if horizontalWin(checks[bnum][x]) || verticalWin(checks[bnum], y) {
						return calculateWin(input, boards[bnum], checks[bnum])
					}
				}
			}
		}
	}
	return 0
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
