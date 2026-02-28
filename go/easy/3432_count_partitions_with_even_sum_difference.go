func countPartitions(nums []int) int {
    totalSum := 0
    for _, v := range nums {
        totalSum += v
    }
    if totalSum % 2 == 0 { // even 4, 
        // 1) even 2+2, even - even = even
        // 2) odd 1+3, odd - odd = even
        // -> all even
        return len(nums) - 1
    }
    // odd:
    // total sum is odd, 4 + 1 = 5, odd + even
    // diff is always odd
    return 0
}

/*
space O(1)
*/
