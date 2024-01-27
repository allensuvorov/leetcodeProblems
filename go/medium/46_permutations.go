func permute(nums []int) [][]int {
    path := make(map[int]bool)
    ans := make([][]int, 0)
    
    var dfs func([]int)
    dfs = func(perm []int) {
        if len(perm) == len(nums) {
            ans = append(ans, slices.Clone(perm))
            return
        }
        if len(perm) > 0 {
            path[perm[len(perm)-1]] = true    
        }
        for _, child := range nums {
            if !path[child] {
                perm = append(perm, child)
                dfs(perm)
                perm = perm[:len(perm)-1]
            }
        }
        if len(perm) > 0 {
            path[perm[len(perm)-1]] = false    
        }
    }
    dfs([]int{})
    return ans
}

