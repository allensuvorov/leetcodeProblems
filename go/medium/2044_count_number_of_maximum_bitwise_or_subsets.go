func countMaxOrSubsets(nums []int) int {
    maxSubset := 0
    for _, v := range nums {
        maxSubset |= v
    }
    // generate sub sets
    return dfs(nums, maxSubset, 0)
}

func dfs(nums []int, maxSubset, sum int) int {
    if len(nums) == 0 {
        return 0
    }

    count := 0
    for i, v := range nums {
        if (sum | v) == maxSubset {
            count++
        }
        count += dfs(nums[i+1:], maxSubset, sum | v)
    }
    
    return count
}
