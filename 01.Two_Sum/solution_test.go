package containerwithmostwater

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		input := []int{2, 7, 11, 15}
		target := 9
		expected := []int{0, 1}
		result := twoSum(input, target)

		if !reflect.DeepEqual(expected, result) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		input := []int{3, 2, 4}
		target := 6
		expected := []int{1, 2}
		result := twoSum(input, target)

		if !reflect.DeepEqual(expected, result) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 3", func(t *testing.T) {
		input := []int{3, 3}
		target := 6
		expected := []int{0, 1}
		result := twoSum(input, target)

		if !reflect.DeepEqual(expected, result) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func twoSum(nums []int, target int) []int {
	prev := make(map[int]int)
	for i, val := range nums {
		k := target - val
		if y, ok := prev[k]; ok {
			return []int{y, i}
		} else {
			prev[val] = i
		}
	}

	return []int{0, 0}
}
