package jumpgame

import (
	"math"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		input := "III"
		expected := 3
		result := romanToInt(input)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		input := "LVIII"
		expected := 58
		result := romanToInt(input)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 3", func(t *testing.T) {
		input := "MCMXCIV"
		expected := 1994
		result := romanToInt(input)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func romanToInt(s string) int {
	dict := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"d": 500,
		"M": 1000,
	}

	prev_num := math.MaxInt32
	res := 0
	for i := 0; i < len(s); i++ {
		c := string(s[i])
		num, exists := dict[c]
		if exists {
			if num > prev_num {
				res = res - prev_num*2 + num
			} else {
				res += num
			}
			prev_num = num
		}
	}

	return res
}
