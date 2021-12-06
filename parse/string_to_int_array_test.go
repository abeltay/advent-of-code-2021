package parse_test

import (
	"reflect"
	"testing"

	"github.com/abeltay/advent-of-code-2021/parse"
)

func TestStringToIntArray(t *testing.T) {
	t.Run("should pass", func(t *testing.T) {
		input := "3,4,3,1,2"
		want := []int{3, 4, 3, 1, 2}
		got, err := parse.StringToIntArray(input)
		if err != nil {
			t.Errorf("got error: %v", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("should have error", func(t *testing.T) {
		input := "3,4,3,abc,2"
		_, err := parse.StringToIntArray(input)
		if err == nil {
			t.Errorf("did not receive error")
		}
	})
}
