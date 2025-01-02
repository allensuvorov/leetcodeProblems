func rob(nums []int) int {
    a, b, c := 0, 0, 0
    for _, v := range nums {
        a, b, c = b, c, v + max(a, b)
    }
    return max(b,c)
}

// Input: nums = [1,2,3,1]
//                  a,b,c
//                1,2,4,3
