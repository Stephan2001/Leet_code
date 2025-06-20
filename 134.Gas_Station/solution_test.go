package productofarrayexceptself

import (
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		gas := []int{1, 2, 3, 4, 5}
		cost := []int{3, 4, 5, 1, 2}
		expected := 3
		result := canCompleteCircuit(gas, cost)

		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		gas := []int{2, 3, 4}
		cost := []int{3, 4, 3}
		expected := -1
		result := canCompleteCircuit(gas, cost)

		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func canCompleteCircuit(gas []int, cost []int) int {
	net_gas := 0
	best_ratio := -100
	best_start := 0
	for i := 0; i < len(gas); i++ {
		net_gas = net_gas - cost[i] + gas[i]
		if gas[i]-cost[i] > best_ratio {
			best_ratio = gas[i] - cost[i]
			best_start = i
		}
	}

	if net_gas >= 0 {
		return best_start
	} else {
		return -1
	}
}
