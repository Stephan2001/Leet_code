package jumpgame

import (
	"sort"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		nums := []int{3, 0, 6, 1, 5}
		expected := 3
		result := hIndex(nums, t)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		nums := []int{1, 3, 1}
		expected := 1
		result := hIndex(nums, t)

		if expected != result {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func hIndex(citations []int, t testing.TB) int {
	citations = Sort(citations)
	count := 0
	for index, val := range citations {
		for k := index; k < len(citations)-1; k++ {
			if (citations[k] - val) >= 0 {
				count++
			} else if count == val {
				return count
			} else {
				break
			}
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
