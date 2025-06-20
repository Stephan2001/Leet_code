package jumpgame

import (
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		input := 3749
		expected := "MMMDCCXLIX"
		result := intToRoman(input)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		input := 58
		expected := "LVIII"
		result := intToRoman(input)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 3", func(t *testing.T) {
		input := 1994
		expected := "MCMXCIV"
		result := intToRoman(input)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

type RomanNumeral struct {
	Value  int
	Symbol string
}

func intToRoman(num int) string {
	numerals := []RomanNumeral{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	res := ""
	total := 0
	for i := 0; total < num; {
		if total+numerals[i].Value <= num {
			total += numerals[i].Value
			res += numerals[i].Symbol
		} else {
			i++
		}
	}

	return res
}
