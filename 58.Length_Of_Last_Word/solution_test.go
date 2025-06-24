package jumpgame

import (
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		input := "Hello World"
		expected := 5
		result := lengthOfLastWord(input)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		input := "   fly me   to   the moon  "
		expected := 4
		result := lengthOfLastWord(input)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 3", func(t *testing.T) {
		input := "luffy is still joyboy"
		expected := 6
		result := lengthOfLastWord(input)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func lengthOfLastWord(s string) int {
	seperated := strings.Fields(s)
	count := 0
	for range seperated[len(seperated)-1] {
		count++
	}
	return count
}
