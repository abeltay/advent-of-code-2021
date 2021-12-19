package aoc

import (
	"fmt"
	"testing"

	"github.com/abeltay/advent-of-code-2021/file"
)

func TestRunner(t *testing.T) {
	/*
		t.Run("test input", func(t *testing.T) {
			data := []string{"[[[[[9,8],1],2],3],4]"}
			want := 143
			got := Runner(data)

			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		})
		t.Run("test input", func(t *testing.T) {
			data := []string{"[7,[6,[5,[4,[3,2]]]]]"}
			want := 143
			got := Runner(data)

			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		})
		t.Run("test input", func(t *testing.T) {
			data := []string{"[[6,[5,[4,[3,2]]]],1]"}
			want := 143
			got := Runner(data)

			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		})
		t.Run("test input", func(t *testing.T) {
			data := []string{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]"}
			want := 143
			got := Runner(data)

			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		})
	*/
	t.Run("test input", func(t *testing.T) {
		data := []string{"[[1,2],[[3,4],5]]"}
		want := 143
		got := Runner(data)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("test input", func(t *testing.T) {
		data := []string{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"}
		want := 1384
		got := Runner(data)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("test input", func(t *testing.T) {
		data := []string{"[[[[1,1],[2,2]],[3,3]],[4,4]]"}
		want := 445
		got := Runner(data)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("test input", func(t *testing.T) {
		data := []string{"[[[[3,0],[5,3]],[4,4]],[5,5]]"}
		want := 791
		got := Runner(data)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("test input", func(t *testing.T) {
		data := []string{"[[[[5,0],[7,4]],[5,5]],[6,6]]"}
		want := 1137
		got := Runner(data)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("test input", func(t *testing.T) {
		data := []string{"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"}
		want := 3488
		got := Runner(data)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("example input", func(t *testing.T) {
		data, err := file.ReadFullData("../example.txt")
		if err != nil {
			t.Fatal(err)
		}
		want := 4140
		got := Runner(data)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("full input", func(t *testing.T) {
		data, err := file.ReadFullData("../input.txt")
		if err != nil {
			t.Fatal(err)
		}
		got := Runner(data)
		fmt.Println(got, "is the answer")
	})
}
