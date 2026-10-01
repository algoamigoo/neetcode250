// my soln

func getConcatenation(nums []int) []int {
    ans :=nums
    for _,val:= range nums{
        ans = append(ans,val)
    }
    return ans
}

-----------------------------
// best soln
func getConcatenation(nums []int) []int {
    ans :=nums
	ans :=append(ans,nums...)
	return ans
}


