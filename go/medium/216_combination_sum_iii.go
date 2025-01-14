func combinationSum3(k int, n int) [][]int {
    res := [][]int{}
    comb := []int{}
    var dfs func(k, sum, num int)
    dfs = func(k, sum, num int) {
        if k == 0 {
            if sum == 0 {
                res = append(res, slices.Clone(comb))
            }
            return
        }

        for num := num; num <=9; num++ {
            comb = append(comb, num)
            dfs(k-1, sum-num, num+1)
            comb = comb[:len(comb)-1]
        }
    }
    dfs(k, n, 1)
    return res
}
