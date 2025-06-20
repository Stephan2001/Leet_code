package productofarrayexceptself

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		expected := []int{24, 12, 8, 6}
		result := productExceptSelf(nums)

		if !reflect.DeepEqual(expected, result) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		nums := []int{-1, 1, 0, -3, 3}
		expected := []int{0, 0, 9, 0, 0}
		result := productExceptSelf(nums)

		if !reflect.DeepEqual(expected, result) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func productExceptSelf(nums []int) []int {
	l := len(nums)
	left := make([]int, l)
	right := make([]int, l)
	ans := make([]int, l)
	left[0] = 1
	right[l-1] = 1
	for i := 1; i < l; i++ {
		j := len(nums) - i - 1
		left[i] = nums[i-1] * left[i-1]
		right[j] = nums[j+1] * right[j+1]
	}

	for i := 0; i < l; i++ {
		ans[i] = right[i] * left[i]
	}
	return ans
}
