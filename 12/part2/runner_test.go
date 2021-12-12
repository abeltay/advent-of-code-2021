package aoc

import (
	"fmt"
	"testing"

	"github.com/abeltay/advent-of-code-2021/file"
)

func TestRunner(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		data := []string{
			"start-A",
			"start-b",
			"A-c",
			"A-b",
			"b-d",
			"A-end",
			"b-end",
		}
		want := 36
		got := Runner(data)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("second input", func(t *testing.T) {
		data := []string{
			"dc-end",
			"HN-start",
			"start-kj",
			"dc-start",
			"dc-HN",
			"LN-dc",
			"HN-end",
			"kj-sa",
			"kj-HN",
			"kj-dc",
		}
		want := 103
		got := Runner(data)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("third input", func(t *testing.T) {
		data := []string{
			"fs-end",
			"he-DX",
			"fs-he",
			"start-DX",
			"pj-DX",
			"end-zg",
			"zg-sl",
			"zg-pj",
			"pj-he",
			"RW-he",
			"fs-DX",
			"pj-RW",
			"zg-RW",
			"start-pj",
			"he-WI",
			"zg-he",
			"pj-fs",
			"start-RW",
		}
		want := 3509
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
