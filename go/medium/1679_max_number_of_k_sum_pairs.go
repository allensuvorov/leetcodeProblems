func maxOperations(nums []int, k int) int {
    counts := map[int]int{}
    res := 0
    for _, num := range nums {   
        diff := k - num
        if counts[diff] > 0 {
            res++
            counts[diff]--
        } else {
            counts[num]++
        }
    }
    return res
}
