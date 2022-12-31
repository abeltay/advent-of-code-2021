package file

import (
	"bufio"
	"fmt"
	"os"
)

// ReadFullData reads the file line by line and outputs as a string array
func ReadFullData(fname string) ([]string, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, fmt.Errorf("Could not open %q: %q", fname, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var data []string
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data, nil
}
