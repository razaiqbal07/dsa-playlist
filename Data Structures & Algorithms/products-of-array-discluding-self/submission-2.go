func productExceptSelf(nums []int) []int {
    result := make([]int,len(nums))
    product := 1
    for i :=0;i<len(nums);i++ {
        result[i] = product
        product = product * nums[i]
    }
    product =1
    for i:=len(nums)-1;i>=0;i-- {
        result[i] =  result[i] * product 
        product = product * nums[i]
    }
    return result
}