func majorityElement(nums []int) int {
    mp:=make(map[int]int)
    for idx:= range nums {
        mp[nums[idx]]++
    }
    for idx:=range mp{
        if mp[idx]>len(nums)/2{
        return idx
        }
    }
    return -1
}

--------------------
func majorityElement(nums []int) int {
    count:=0
    curr:=nums[0] 
    for _,val:= range nums {
        if curr==val {
            count++
        } else {
            count--
        }
        if count<0 {
            curr = val
            count = 1
        }
    }
    return curr
}