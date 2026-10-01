func removeElement(nums []int, val int) int {
	i := 0
	n := len(nums)
	j := n - 1
	count := 0
	for _, value := range nums {
		if value == val {
			count++
		}
	}
	for i <= j {
		for i <= j && nums[j] == val {
			j--
		}
		if i <= j && nums[i] == val {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}
		i++
	}
	return n - count

}