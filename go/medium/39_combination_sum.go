func combinationSum(candidates []int, target int) [][]int {
    ans := [][]int{}
    comb := []int{}

    var dfs func(sum int, curCandidates []int)
    dfs = func(sum int, curCandidates []int) {
        if sum == target {
            ans = append(ans, slices.Clone(comb))
            return
        }
        for i, v := range curCandidates {
            if target - sum >= v {
                comb = append(comb, v)
                dfs(sum + v, curCandidates[i:])
                comb = comb[:len(comb) - 1]
            }
        }
    }
    dfs(0, candidates)
    return ans
}
