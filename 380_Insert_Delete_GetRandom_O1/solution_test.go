package jumpgame

import (
	"math/rand"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		obj := Constructor()

		if !obj.Insert(10) {
			t.Errorf("Expected Insert(10) to be true")
		}

		if obj.Insert(10) {
			t.Errorf("Expected Insert(10) second time to be false")
		}

		if !obj.Remove(10) {
			t.Errorf("Expected Remove(10) to be true")
		}

		if obj.Remove(20) {
			t.Errorf("Expected Remove(20) to be false")
		}

		obj.Insert(1)
		obj.Insert(2)
		obj.Insert(3)

		val := obj.GetRandom()
		if val != 1 && val != 2 && val != 3 {
			t.Errorf("GetRandom() returned unexpected value %d", val)
		}
	})
}

type RandomizedSet struct {
	arr  []int
	size int
	set  map[int]int
}

func Constructor() RandomizedSet {
	arr := make([]int, 0)
	size := 0
	set := make(map[int]int)
	return RandomizedSet{
		arr:  arr,
		size: size,
		set:  set,
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.set[val]
	if ok {
		return false
	}

	this.arr = append(this.arr, val)
	this.set[val] = this.size
	this.size++
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.set[val]
	if !ok {
		return false
	}

	lastElement := this.arr[this.size-1]
	this.arr[index] = lastElement
	this.set[lastElement] = index

	this.arr = this.arr[:this.size-1]
	delete(this.set, val)
	this.size--
	return true
}

func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(this.size)
	return this.arr[index]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
