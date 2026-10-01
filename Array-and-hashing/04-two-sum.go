// my soln
func twoSum(nums []int, target int) []int {

	mp := make(map[int]int)
	ans := []int{}
	for idx := range nums {
		// check if we have curr - target in our map
		val, ok := mp[target-nums[idx]]

		if ok == true {
			ans = append(ans, idx, val)
			return ans
		}
		// add the idx of the nums element in the map
		mp[nums[idx]] = idx
	}
	return ans
}

//  best soln

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		if k, ok := m[v]; ok {
			return []int{k, i}
		}
		m[target-v] = i
	}
	return []int{}
}