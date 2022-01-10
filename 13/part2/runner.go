package aoc

import (
	"log"
	"strconv"
	"strings"

	"github.com/abeltay/advent-of-code-2021/parse"
)

type fold struct {
	vertical bool
	line     int
}

// Runner runs the algorithm to get the answer
func Runner(data []string) [][]int {
	var maxX, maxY int
	var parseFold bool
	var folds []fold
	for i := range data {
		if data[i] == "" {
			parseFold = true
			continue
		}
		if parseFold {
			s := strings.Split(data[i], " ")
			s1 := strings.Split(s[2], "=")
			num, err := strconv.Atoi(s1[1])
			if err != nil {
				log.Fatal(err)
			}
			if s1[0] == "y" {
				folds = append(folds, fold{line: num})
			} else {
				folds = append(folds, fold{vertical: true, line: num})
			}
			continue
		}
		arr, err := parse.StringToIntArray(data[i])
		if err != nil {
			log.Fatal(err)
		}
		x, y := arr[0], arr[1]
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}
	// fmt.Println(maxX, maxY)
	img := make([][]int, maxY+1)
	for i := range img {
		img[i] = make([]int, maxX+1)
	}
	for i := range data {
		if data[i] == "" {
			break
		}
		arr, err := parse.StringToIntArray(data[i])
		if err != nil {
			log.Fatal(err)
		}
		x, y := arr[0], arr[1]
		img[y][x]++
		// fmt.Printf("%d %d\n", x, y)
	}
	/*
		for i := range img {
			fmt.Printf("%v\n", img[i])
		}
	*/
	for n := range folds {
		if !folds[n].vertical {
			for i := 0; i < folds[n].line; i++ {
				if folds[n].line+1+i >= len(img) || folds[n].line-1-i < 0 {
					continue
				}
				for j := range img[i] {
					img[folds[n].line-1-i][j] += img[folds[n].line+1+i][j]
				}
			}
			img = img[:folds[n].line]
		} else {
			for i := range img {
				for j := 0; j < folds[n].line; j++ {
					img[i][folds[n].line-1-j] += img[i][folds[n].line+1+j]
				}
			}
			for i := range img {
				img[i] = img[i][:folds[n].line]
			}
		}
	}
	return img
}
