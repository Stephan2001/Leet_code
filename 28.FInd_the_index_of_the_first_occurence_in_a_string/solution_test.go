package findtheindexofthefirstoccurenceinastring

import "testing"

func TestStrStr(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		input := "sadbutsad"
		heystack := "sad"
		expected := 0
		result := strStr(input, heystack)

		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		input := "leetcode"
		heystack := "leeto"
		expected := -1
		result := strStr(input, heystack)

		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("Case 3", func(t *testing.T) {
		input := "Beeeeo"
		heystack := "oo"
		expected := -1
		result := strStr(input, heystack)

		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func strStr(haystack string, needle string) int {
	l_needle := len(needle)
	i, k := 0, 0
	for ; i < len(haystack) && k < l_needle; i++ {
		if needle[k] == haystack[i] {
			k++
		} else {
			k = 0
		}
	}
	if k < l_needle {
		return -1
	}
	return i - k
}
