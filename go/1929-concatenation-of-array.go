func getConcatenation(nums []int) []int {
    n := len(nums)
    res := make([]int, n*2)
    for i, v := range nums {
        res[i] = v
        res[i+n] = v
    }
    return res
}
