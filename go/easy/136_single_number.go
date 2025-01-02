func singleNumber(nums []int) int {
    num := 0
    for _, v := range nums {
        num ^= v
    }
    return num
}
