package besttimetobuyandsellstock

import (
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		nums := []int{7, 1, 5, 3, 6, 4}
		expected := 5
		result := maxProfit(nums)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		nums := []int{7, 6, 4, 3, 1}
		expected := 0
		result := maxProfit(nums)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func maxProfit(prices []int) int {
	profit := 0
	for i := 0; i < len(prices); i++ {
		for k := i + 1; k < len(prices); k++ {
			if prices[k]-prices[i] > profit {
				profit = prices[k] - prices[i]
			}
		}
	}
	return profit
}
