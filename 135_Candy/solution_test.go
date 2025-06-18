package majorityelement

import (
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		nums := []int{1, 0, 2}
		expected := 5
		result := candy(nums, t)

		if expected != result {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		nums := []int{1, 2, 2}
		expected := 4
		result := candy(nums, t)

		if expected != result {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
	t.Run("Case 3", func(t *testing.T) {
		nums := []int{1, 2, 87, 4, 3, 2, 1}
		expected := 18
		result := candy(nums, t)

		if expected != result {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}

func candy(ratings []int, t testing.TB) int {
	len := len(ratings)
	result := 0
	for i := 0; i < len-1; i++ {
		if ratings[i] > ratings[i+1] {
			result++
		}
		if ratings[len-i-1] > ratings[len-i-2] {
			result++
		}
	}
	return result + len
}
