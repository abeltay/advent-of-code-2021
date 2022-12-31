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
			"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1",
			"",
			"22 13 17 11  0",
			" 8  2 23  4 24",
			"21  9 14 16  7",
			" 6 10  3 18  5",
			" 1 12 20 15 19",
			"",
			" 3 15  0  2 22",
			" 9 18 13 17  5",
			"19  8  7 25 23",
			"20 11 10 24  4",
			"14 21 16 12  6",
			"",
			"14 21 17 24  4",
			"10 16 15  9 19",
			"18  8 23 26 20",
			"22 11 13  6  5",
			" 2  0 12  3  7",
		}
		want := 4512
		got := Runner(data)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("vertical win input", func(t *testing.T) {
		data := []string{
			"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1",
			"",
			"22 13 17  3 20",
			" 8  2 23  4 24",
			"21  9 14 16  7",
			" 6 10 11 18  5",
			" 1 12  0 15 19",
			"",
			" 3 15  0  2 22",
			" 9 18 13 17  5",
			"19  8  7 25 23",
			"20 11 10 24  4",
			"14 21 16 12  6",
			"",
			"14 21 17 24  4",
			"10 16 15  9 19",
			"18  8 23 26 20",
			"22 11 13  6  5",
			" 2  0 12  3  7",
		}
		want := 2912
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
