package removeDuplicates
func removeDuplicates(nums []int) int {
    seen := make(map[int]bool)
	k := 0

	for i := 0; i < len(nums); i++ {
		if !seen[nums[i]] {
			seen[nums[i]] = true
			nums[k] = nums[i]
			k++
		}
	}

	return k
}