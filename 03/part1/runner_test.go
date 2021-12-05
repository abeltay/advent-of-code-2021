package part1

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestRunner(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		data := []string{
			"00100",
			"11110",
			"10110",
			"10111",
			"10101",
			"01111",
			"00111",
			"11100",
			"10000",
			"11001",
			"00010",
			"01010",
		}
		want := 198
		got := Runner(data)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("full input", func(t *testing.T) {
		data := readFullData(t, "../input.txt")
		got := Runner(data)
		fmt.Println(got, "is the answer")
	})
}

func readFullData(t *testing.T, fname string) []string {
	t.Helper()
	f, err := os.Open(fname)
	if err != nil {
		t.Fatalf("Could not open %q: %q", fname, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var data []string
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}
