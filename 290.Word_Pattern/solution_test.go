package issubsequence

import (
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		input := "abba"
		pattern := "dog cat cat dog"
		expected := true
		result := wordPattern(input, pattern)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		input := "abba"
		pattern := "dog cat cat fish"
		expected := false
		result := wordPattern(input, pattern)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 3", func(t *testing.T) {
		input := "aaaa"
		pattern := "dog cat cat dog"
		expected := false
		result := wordPattern(input, pattern)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func wordPattern(pattern string, s string) bool {
	letters := map[byte]int{}
	words := map[string]int{}

	word := ""
	index := 0
	for i := 0; i < len(s); i++ {
		val := s[i]
		if val != ' ' {
			word = word + string(val)
			if i != len(s)-1 {
				continue
			}
		}

		if _, ok := words[word]; ok {
			words[word]++
		} else {
			words[word] = 0
		}
		word = ""

		c := pattern[index]
		if _, ok := letters[c]; ok {
			letters[c]++
		} else {
			letters[c] = 0
		}
		index++

		if len(letters) != len(words) {
			return false
		}
	}

	return true
}
