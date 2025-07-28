package productofarrayexceptself

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		input_a := "anagram🧠"
		input_b := "nagaram🧠"
		expected := true
		result := isAnagram(input_a, input_b)

		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		input_a := "rat"
		input_b := "car"
		expected := false
		result := isAnagram(input_a, input_b)

		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func isAnagram(s string, t string) bool {
	map_a := map[byte]int{}
	map_b := map[byte]int{}

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		val_a := s[i]
		if _, ok := map_a[val_a]; ok {
			map_a[val_a]++
		} else {
			map_a[val_a] = 0
		}

		val_b := t[i]
		if _, ok := map_b[val_b]; ok {
			map_b[val_b]++
		} else {
			map_b[val_b] = 0
		}
	}

	return reflect.DeepEqual(map_a, map_b)
}
