func maxOperations(nums []int, k int) int {
    counts := map[int]int{}
    res := 0
    for _, v := range nums {
        if counts[k-v] > 0 {
            res++
            counts[k-v]--
        } else {
            counts[v]++
        }
    }
    return res
}
