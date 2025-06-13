package jumpgame

import (
	"sort"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		nums := []int{3, 0, 6, 1, 5}
		expected := 3
		result := hIndex(nums)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		nums := []int{1, 3, 1}
		expected := 1
		result := hIndex(nums)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 3", func(t *testing.T) {
		nums := []int{2, 2, 5}
		expected := 2
		result := hIndex(nums)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func hIndex(citations []int) int {
	citations = Sort(citations)
	for i := 0; i < len(citations); i++ {
		if citations[i] <= i+1 {
			return citations[i]
		}
	}
	return 0
}

func Sort(citations []int) []int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})
	return citations
}
