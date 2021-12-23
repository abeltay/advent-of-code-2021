package aoc

import (
	"fmt"
	"testing"

	"github.com/abeltay/advent-of-code-2021/file"
)

func TestRunner(t *testing.T) {
	t.Run("test input", func(t *testing.T) {
		data := []string{
			"on x=10..12,y=10..12,z=10..12",
			"on x=11..13,y=11..13,z=11..13",
			"off x=9..11,y=9..11,z=9..11",
			"on x=10..10,y=10..10,z=10..10",
		}
		want := 39
		got := Runner(data)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("example input", func(t *testing.T) {
		data, err := file.ReadFullData("example.txt")
		if err != nil {
			t.Fatal(err)
		}
		want := 590784
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
