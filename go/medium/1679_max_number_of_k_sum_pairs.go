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

// practice 2
func maxOperations(nums []int, k int) int {
    counts := map[int]int{}
    for _, v := range nums {
        counts[v]++
    }
    res := 0
    for num, count := range counts {
        numOps := 0
        if num * 2 == k {
            numOps = count/2
            counts[num] = 0
        } else {
            numOps = min(count, counts[k-num])
            counts[num] = 0
            counts[k-num] = 0
        }
        res += numOps
    }
    return res
}
