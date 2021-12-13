package aoc

import (
	"fmt"
	"testing"

	"github.com/abeltay/advent-of-code-2021/file"
)

func TestRunner(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		data, err := file.ReadFullData("../example.txt")
		if err != nil {
			t.Fatal(err)
		}
		want := 10
		got := Runner(data)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("second input", func(t *testing.T) {
		data, err := file.ReadFullData("../example2.txt")
		if err != nil {
			t.Fatal(err)
		}
		want := 19
		got := Runner(data)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("third input", func(t *testing.T) {
		data, err := file.ReadFullData("../example3.txt")
		if err != nil {
			t.Fatal(err)
		}
		want := 226
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
