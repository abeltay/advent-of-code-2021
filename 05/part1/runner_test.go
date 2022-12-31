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
			"0,9 -> 5,9",
			"8,0 -> 0,8",
			"9,4 -> 3,4",
			"2,2 -> 2,1",
			"7,0 -> 7,4",
			"6,4 -> 2,0",
			"0,9 -> 2,9",
			"3,4 -> 1,4",
			"0,0 -> 8,8",
			"5,5 -> 8,2",
		}
		want := 5
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
