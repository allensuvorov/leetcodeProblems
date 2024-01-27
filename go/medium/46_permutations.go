func permute(nums []int) [][]int {
    m := make(map[int]bool)
    ans := make([][]int, 0)
    for _, v := range nums {
        m[v] = true
    }
    var dfs func(int, map[int]bool, []int)
    dfs = func(now int, rem map[int]bool, perm [] int) {
        if len(perm) == len(nums) {
            tmp := slices.Clone(perm)
            ans = append(ans, tmp)
            return
        }

        for next, ok := range rem {
            if ok {
                perm = append(perm, next)
                rem[next] = false
                dfs(next, rem, perm)
                rem[next] = true
                perm = perm[:len(perm)-1]
            }
        }
    }
    dfs(0, m, []int{})
    return ans
}

