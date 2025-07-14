package issubsequence

import (
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		input_s := "abd"
		input_t := "ahbgdc"
		expected := true
		result := isSubsequence(input_s, input_t)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		input_s := "axc"
		input_t := "ahbgdc"
		expected := false
		result := isSubsequence(input_s, input_t)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func isSubsequence(s string, t string) bool {
	l_s, l_t := len(s), len(t)
	for i, c := 0, 0; i < l_t; i++ {
		if t[i] == s[c] {
			c++
		}
		if l_s == c {
			return true
		}
	}
	return false
}
