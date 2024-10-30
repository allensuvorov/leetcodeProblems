func combinationSum3(k int, n int) [][]int {
    ans := [][]int{}
    nums := []int{1,2,3,4,5,6,7,8,9}
    
    var dfs func(nums, comb []int, sum int)
    dfs = func(nums, comb []int, sum int) {
        if len(comb) == k && sum == n {
            ans = append(ans, slices.Clone(comb))
            return
        }

        if len(comb) > k || sum > n {
            return
        }
        
        for i, v := range nums {
            dfs(nums[i+1:], append(comb, v), sum + v)
        }
    }
    
    dfs(nums, []int{}, 0)
    return ans
}
