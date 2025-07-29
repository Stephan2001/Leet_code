package jumpgame

import (
	"sort"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
		expected := [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}
		result := groupAnagrams(input)

		if !compareGroups(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		input := []string{""}
		expected := [][]string{{""}}
		result := groupAnagrams(input)

		if !compareGroups(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 3", func(t *testing.T) {
		input := []string{"a"}
		expected := [][]string{{"a"}}
		result := groupAnagrams(input)

		if !compareGroups(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func normalizeGroup(group []string) string {
	sort.Strings(group)
	return strings.Join(group, ",")
}

func compareGroups(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	normalize := func(groups [][]string) map[string]bool {
		m := make(map[string]bool)
		for _, group := range groups {
			key := normalizeGroup(group)
			m[key] = true
		}
		return m
	}

	amap := normalize(a)
	bmap := normalize(b)

	if len(amap) != len(bmap) {
		return false
	}

	for k := range amap {
		if !bmap[k] {
			return false
		}
	}
	return true
}

func groupAnagrams(strs []string) [][]string {
	list := make(map[string][]string)
	for _, val := range strs {
		runes := []rune(val)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		k := string(runes)
		list[k] = append(list[k], val)
	}
	res := [][]string{}
	for _, group := range list {
		res = append(res, group)
	}
	return res
}
