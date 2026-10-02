func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	left := sortArray(nums[:len(nums)/2])
	right := sortArray(nums[len(nums)/2:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := []int{}

	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for ; i < len(left); i++ {
		result = append(result, left[i])
	}
	for ; j < len(right); j++ {
		result = append(result, right[j])
	}

	return result
}