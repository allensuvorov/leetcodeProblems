func minimumOperations(nums []int) int {
    seen := make(map[int]bool)
    i := len(nums) - 1; 
    for i >= 0 && !seen[nums[i]] {
        seen[nums[i]] = true
        i--
    }

    if i == -1 {
        return 0
    }

    k := i + 1
    res := k / 3
    if k % 3 != 0 {
        res++
    }
    return res
}
