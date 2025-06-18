package besttimetobuyandsellstock

import (
	"math"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		nums := []int{7, 1, 5, 3, 6, 4}
		expected := 7
		result := maxProfit(nums, t)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		expected := 4
		result := maxProfit(nums, t)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		nums := []int{7, 6, 4, 3, 1}
		expected := 0
		result := maxProfit(nums, t)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func Max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func maxProfit(prices []int, t testing.TB) int {

	cur_hold, cur_not_hold := math.MinInt64, 0

	for _, stock_price := range prices {
		t.Logf("Current price: %d", stock_price)
		t.Logf("prev keep: %d", cur_hold)
		t.Logf("prev yeet: %d", cur_not_hold)
		prev_hold, prev_not_hold := cur_hold, cur_not_hold

		cur_hold = Max(prev_hold, prev_not_hold-stock_price)
		t.Logf("current keep: %d", cur_hold)
		cur_not_hold = Max(prev_not_hold, prev_hold+stock_price)
		t.Logf("current yeet: %d", cur_not_hold)
		t.Log("=============")
	}

	return cur_not_hold

}
