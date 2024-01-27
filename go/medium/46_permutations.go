func permute(nums []int) [][]int {
    m := make(map[int]bool)
    ans := make([][]int, 0)
    for _, v := range nums {
        m[v] = true
    }
    var dfs func(int, map[int]bool, []int)
    dfs = func(now int, rem map[int]bool, perm [] int) {
        if len(rem) == 0 {
            ans = append(ans, perm)
            return
        }

        for next := range rem {
            newPerm := slices.Clone(perm)
            newPerm = append(newPerm, next)

            newRem := make(map[int]bool)
            for k := range rem {
                if k != next {
                    newRem[k] = true
                }
            }
            dfs(next, newRem, newPerm)
        }
    }
    dfs(0, m, []int{})
    return ans
}

