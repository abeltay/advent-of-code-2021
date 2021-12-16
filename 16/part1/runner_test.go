package aoc

import (
	"fmt"
	"testing"

	"github.com/abeltay/advent-of-code-2021/file"
)

func TestRunner(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		data := []string{"D2FE28"}
		want := 6
		got := Runner(data)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("example input", func(t *testing.T) {
		data := []string{"38006F45291200"}
		want := 9
		got := Runner(data)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("example input", func(t *testing.T) {
		data := []string{"EE00D40C823060"}
		want := 14
		got := Runner(data)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("example input", func(t *testing.T) {
		data := []string{"8A004A801A8002F478"}
		want := 16
		got := Runner(data)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("example input", func(t *testing.T) {
		data := []string{"620080001611562C8802118E34"}
		want := 12
		got := Runner(data)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("example input", func(t *testing.T) {
		data := []string{"C0015000016115A2E0802F182340"}
		want := 23
		got := Runner(data)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("example input", func(t *testing.T) {
		data := []string{"A0016C880162017C3686B18A3D4780"}
		want := 31
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
