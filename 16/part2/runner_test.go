package aoc

import (
	"fmt"
	"testing"

	"github.com/abeltay/advent-of-code-2021/file"
)

func TestRunner(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		data := []string{"C200B40A82"}
		want := 3
		got := Runner(data)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("example input", func(t *testing.T) {
		data := []string{"04005AC33890"}
		want := 54
		got := Runner(data)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("example input", func(t *testing.T) {
		data := []string{"880086C3E88112"}
		want := 7
		got := Runner(data)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("example input", func(t *testing.T) {
		data := []string{"CE00C43D881120"}
		want := 9
		got := Runner(data)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("example input", func(t *testing.T) {
		data := []string{"D8005AC2A8F0"}
		want := 1
		got := Runner(data)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("example input", func(t *testing.T) {
		data := []string{"F600BC2D8F"}
		want := 0
		got := Runner(data)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("example input", func(t *testing.T) {
		data := []string{"9C005AC2F8F0"}
		want := 0
		got := Runner(data)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("example input", func(t *testing.T) {
		data := []string{"9C0141080250320F1802104A08"}
		want := 1
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
