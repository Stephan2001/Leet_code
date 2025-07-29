package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		input := []int{0, 1, 2, 4, 5, 7}
		expected := []string{"0->2", "4->5", "7"}
		result := summaryRanges(input)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		input := []int{0, 2, 3, 4, 6, 8, 9}
		expected := []string{"0", "2->4", "6", "8->9"}
		result := summaryRanges(input)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func summaryRanges(nums []int) []string {
	l := len(nums)

	if l == 0 {
		return []string{}
	}
	result := []string{}
	start := 0
	for i := 1; i <= l; i++ {
		if i == l || nums[i] != nums[i-1]+1 {
			if nums[start] == nums[i-1] {
				result = append(result, fmt.Sprintf("%d", nums[start]))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", nums[start], nums[i-1]))
			}
			if i < l {
				start = i
			}
		}
	}

	return result
}
