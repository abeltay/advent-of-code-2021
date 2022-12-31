package aoc

import (
	"testing"

	"github.com/abeltay/advent-of-code-2021/file"
)

func TestRunner(t *testing.T) {
	t.Run("full input", func(t *testing.T) {
		data, err := file.ReadFullData("input.txt")
		if err != nil {
			t.Fatal(err)
		}
		Runner(data)
	})
}
