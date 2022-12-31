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

func TestContinuousStringToIntArray(t *testing.T) {
	t.Run("should pass", func(t *testing.T) {
		input := "2199943210"
		want := []int{2, 1, 9, 9, 9, 4, 3, 2, 1, 0}
		got, err := parse.ContinuousStringToIntArray(input)
		if err != nil {
			t.Errorf("got error: %v", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("should have error", func(t *testing.T) {
		input := "21999432a0"
		_, err := parse.ContinuousStringToIntArray(input)
		if err == nil {
			t.Errorf("did not receive error")
		}
	})
}
