func productExceptSelf(nums []int) []int {
    products := make([]int,len(nums))
    for i,_ := range nums {
        product := 1
        for j,num1 := range nums {
            if i != j {
                product = product * num1
            }
        }
        products[i] = product
    }
    return products
}
