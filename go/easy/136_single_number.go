func singleNumber(nums []int) int {
    single := 0 
    for _, v := range nums {
        single = single ^ v
    }
    return single
}
