package twosumiiinputarrayissorted

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		nums := []int{2, 7, 11, 15}
		target := 9
		expected := []int{1, 2}
		result := twoSum(nums, target)

		if reflect.DeepEqual(result, target) {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		nums := []int{2, 3, 4}
		expected := []int{1, 3}
		target := 6
		result := twoSum(nums, target)

		if reflect.DeepEqual(result, target) {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
	t.Run("Case 3", func(t *testing.T) {
		nums := []int{-1, 0}
		expected := []int{1, 2}
		target := -1
		result := twoSum(nums, target)

		if reflect.DeepEqual(result, target) {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}

func twoSum(numbers []int, target int) []int {
	a, b := 0, len(numbers)-1
	for range numbers {
		sum := numbers[a] + numbers[b]

		if sum == target {
			return []int{a + 1, b + 1}
		}

		if sum > target {
			b--
		} else {
			a++
		}
	}

	return []int{}
}
