package rotatearray

import (
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		input_s := "egg"
		input_t := "add"
		expected := true
		result := isIsomorphic(input_s, input_t)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		input_s := "foo"
		input_t := "bar"
		expected := false
		result := isIsomorphic(input_s, input_t)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 3", func(t *testing.T) {
		input_s := "paper"
		input_t := "title"
		expected := true
		result := isIsomorphic(input_s, input_t)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func isIsomorphic(s string, t string) bool {
	arr_s := map[byte]int{}
	arr_t := map[byte]int{}

	for i := 0; i < len(s); i++ {
		val := s[i]
		if _, ok := arr_s[val]; ok {
			arr_s[val]++
		} else {
			arr_s[val] = 0
		}
		val = t[i]
		if _, ok := arr_t[val]; ok {
			arr_t[val]++
		} else {
			arr_t[val] = 0
		}

		if len(arr_s) != len(arr_t) {
			return false
		}
	}

	return true
}
