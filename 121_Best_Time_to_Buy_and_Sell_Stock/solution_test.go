package besttimetobuyandsellstock

import (
	"math"
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
	max_sell := 0
	min_buy := math.MaxInt32

	for _, price := range prices {
		if price < min_buy {
			min_buy = price
		} else if (price - min_buy) > max_sell {
			max_sell = price
		}
	}

	return max_sell
}
