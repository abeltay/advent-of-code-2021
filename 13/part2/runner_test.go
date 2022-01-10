package aoc

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/abeltay/advent-of-code-2021/file"
)

func TestRunner(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		data, err := file.ReadFullData("../example.txt")
		if err != nil {
			t.Fatal(err)
		}
		want := [][]int{
			{1, 1, 1, 1, 1},
			{1, 0, 0, 0, 1},
			{1, 0, 0, 0, 1},
			{1, 0, 0, 0, 1},
			{1, 2, 2, 1, 1},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		}
		got := Runner(data)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("full input", func(t *testing.T) {
		data, err := file.ReadFullData("../input.txt")
		if err != nil {
			t.Fatal(err)
		}
		got := Runner(data)
		for i := range got {
			for j := range got[i] {
				if got[i][j] == 0 {
					fmt.Print(" ")
				} else {
					fmt.Print("█")
				}
			}
			fmt.Println()
		}
	})
}
