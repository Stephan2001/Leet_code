package rotatearray

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6, 7}
		expected := []int{5, 6, 7, 1, 2, 3, 4}
		result := rotate(nums, 3)

		if !reflect.DeepEqual(expected, result) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		nums := []int{-1, -100, 3, 99}
		expected := []int{3, 99, -1, -100}
		result := rotate(nums, 2)

		if !reflect.DeepEqual(expected, result) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func rotate(nums []int, k int) []int {
	b := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		b[(i+k)%len(nums)] = nums[i]
	}
	nums = b
	return nums
}
