package jumpgame

import (
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		input := "the sky is blue"
		expected := "blue is sky the"
		result := reverseWords(input)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		input := "  hello world  "
		expected := "world hello"
		result := reverseWords(input)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 3", func(t *testing.T) {
		input := "a good   example"
		expected := "example good a"
		result := reverseWords(input)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func reverseWords(s string) string {
	seperated := strings.Fields(s)
	l := len(seperated)
	for i, k := 0, l-1; i < l/2; {
		temp := seperated[i]
		seperated[i] = seperated[k]
		seperated[k] = temp
		i++
		k--
	}
	return strings.Join(seperated, " ")
}
