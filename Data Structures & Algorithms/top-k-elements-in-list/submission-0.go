func topKFrequent(nums []int, k int) []int {
    result := []int{}
    frequencyBucket := make([][]int,len(nums) + 1)
    frequencyMap := map[int]int{}
    for i:=0;i<len(nums);i++ {
        frequencyMap[nums[i]] = frequencyMap[nums[i]] + 1
    }
    for key,value := range frequencyMap {
        frequencyBucket[value] = append(frequencyBucket[value],key)
    }
    for i := len(frequencyBucket) - 1;i>0;i-- {
        result =append(result,frequencyBucket[i]...)
        if len(result) >= k {return result[:k]}
    }
    return result
}
