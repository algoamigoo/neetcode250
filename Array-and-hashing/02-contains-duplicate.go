func containsDuplicate(nums []int) bool {

    mp:=make(map[int]int)
    for idx,_ :=range nums{
        mp[nums[idx]]++
    }


    for _,val := range mp {
        if(val>1){
        return true
        }
    }
    
    return false
}