package ransomnote

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		input_a := "a"
		input_b := "b"
		expected := false
		result := canConstruct(input_a, input_b)

		if !reflect.DeepEqual(expected, result) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		input_a := "aa"
		input_b := "ab"
		expected := false
		result := canConstruct(input_a, input_b)

		if !reflect.DeepEqual(expected, result) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		input_a := "aa"
		input_b := "aab"
		expected := true
		result := canConstruct(input_a, input_b)

		if !reflect.DeepEqual(expected, result) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func canConstruct(ransomNote string, magazine string) bool {
	newmap := map[rune]int{}

	for _, char := range magazine {
		newmap[char]++
	}

	for _, char := range ransomNote {
		if _, ok := newmap[char]; ok {
			newmap[char]--
		} else {
			return false
		}
	}

	for _, val := range newmap {
		if val < 0 {
			return false
		}
	}

	return true
}
