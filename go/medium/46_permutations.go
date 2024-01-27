func permute(nums []int) [][]int {
    m := make(map[int]bool)
    ans := make([][]int, 0)
    for _, v := range nums {
        m[v] = true
    }
    var dfs func(num int, rem map[int]bool)
    dfs = func(now int, rem map[int]bool) {
        if len(rem) == 0 {
            newPerm := make([]int, 0, len(nums))
            newPerm = append(newPerm, now)
            ans = append(ans, newPerm)
            return
        }

        for next := range rem {
            newRem := make(map[int]bool)
            for k := range rem {
                if k != next {
                    newRem[k] = true
                }
            }
            dfs(next, newRem)

            i := len(ans)-1
            // if len(ans[i]) < len(nums) {
            //     ans[i]= append(ans[i], next)
            // }
            if len(ans[i]) < len(nums) {
                ans[i]= append(ans[i], now)
            }
        }
    }
    dfs(0, m)
    return ans
}

