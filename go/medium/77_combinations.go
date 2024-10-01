// O(K∗C(N,K))

func combine(n int, k int) [][]int {
    res := [][]int{}
    comb := []int{}
    var dfs func(cur int)
    dfs = func(cur int) {
        if len(comb) == k {
            res = append(res, slices.Clone(comb))
            return
        }
        for num := cur; num <= n; num++ {
            comb = append(comb, num)
            dfs(num+1)
            comb = comb[:len(comb)-1]
        }
    }
    dfs(1)
    return res
}
