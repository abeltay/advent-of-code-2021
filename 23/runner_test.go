package aoc

import (
	"fmt"
	"testing"
)

func TestRunner(t *testing.T) {
	t.Run("part 1 example input", func(t *testing.T) {
		data := space{
			hallway: []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
			rooms: [][]int{
				{0, 1},
				{3, 2},
				{2, 1},
				{0, 3},
			},
		}
		want := 12521
		got := Runner(data)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("part 1 full input", func(t *testing.T) {
		data := space{
			hallway: []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
			rooms: [][]int{
				{1, 3},
				{0, 0},
				{3, 1},
				{2, 2},
			},
		}
		got := Runner(data)
		fmt.Println(got, "is the answer")
	})
	t.Run("part 2 example input", func(t *testing.T) {
		data := space{
			hallway: []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
			rooms: [][]int{
				{0, 3, 3, 1},
				{3, 1, 2, 2},
				{2, 0, 1, 1},
				{0, 2, 0, 3},
			},
		}
		want := 44169
		got := Runner(data)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("part 2 full input", func(t *testing.T) {
		data := space{
			hallway: []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
			rooms: [][]int{
				{1, 3, 3, 3},
				{0, 1, 2, 0},
				{3, 0, 1, 1},
				{2, 2, 0, 2},
			},
		}
		got := Runner(data)
		fmt.Println(got, "is the answer")
	})
}
