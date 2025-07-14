package besttimetobuyandsellstock

import (
	"regexp"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		input := "Ama"
		expected := true
		result := isPalindrome(input)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		input := "A man, a plan, a canal: Panama"
		expected := true
		result := isPalindrome(input)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 3", func(t *testing.T) {
		input := "race a car"
		expected := false
		result := isPalindrome(input)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 4", func(t *testing.T) {
		input := " "
		expected := true
		result := isPalindrome(input)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	regx := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	s = regx.ReplaceAllString(s, "")

	l := len(s)

	if l == 1 || l == 0 {
		return true
	}

	for i := 0; i < (l / 2); i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}

	return true
}
