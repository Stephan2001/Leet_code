package containerwithmostwater

import (
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
		expected := 49
		result := maxArea(input)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		input := []int{1, 1}
		expected := 1
		result := maxArea(input)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func Max(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Min(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxArea(height []int) int {
	l := len(height)
	a, b := 0, l-1
	current_max := 0
	for a < b {
		current_width := b - a
		current_Height := min(height[a], height[b])
		current_max = max(current_max, current_Height*current_width)

		if height[a] <= height[b] {
			a++
		} else {
			b--
		}
	}

	return current_max
}
