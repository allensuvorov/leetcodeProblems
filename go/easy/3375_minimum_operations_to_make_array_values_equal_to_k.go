func minOperations(nums []int, k int) int {
    set := make([]bool, 101)

    for _, v := range nums {
        set[v] = true
    }

    res := 0
    for i := 1; i < len(set); i++ {
        if set[i] {
            if i < k {
                return -1
            }
            if i > k {
                res++
            }
        }
    }

    return res
} 
