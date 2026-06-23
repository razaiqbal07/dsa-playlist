func hasDuplicate(nums []int) bool {
    traversed := make(map[int]bool)
    for i:=0;i<len(nums);i++ {
        _,ok := traversed[nums[i]]
        if ok {return true}
        traversed[nums[i]] = true
    }
    return false
}
