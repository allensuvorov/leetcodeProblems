func permute(nums []int) [][]int {
    ans := make([][]int, 0)
    
    if len(nums) == 1 {
        return append(ans, nums)
    }
    
    for _ = range nums {
        n := nums[0]
        nums = nums[1:]
        perms := permute(nums)
        for i := range perms {
            perms[i] = append(perms[i], n)
        }
        ans = append(ans, perms...)
        nums = append(nums, n)
    }
    return ans
}
