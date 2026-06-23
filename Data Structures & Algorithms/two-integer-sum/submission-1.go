func twoSum(nums []int, target int) []int {
    numsMap := map[int]int {}
    for i:=0;i<len(nums);i++ {
        required := target - nums[i]
        value,ok := numsMap[required]
        if ok {
            return []int{value,i}
        }else {
            numsMap[nums[i]] = i
        }
    }
    return nil
}
