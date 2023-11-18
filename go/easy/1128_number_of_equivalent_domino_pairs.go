func numEquivDominoPairs(dominoes [][]int) int {
    res := 0
    dm := make(map[[10]int]int)
    for _, d := range dominoes {
        na := [10]int{}
        na[d[0]]++
        na[d[1]]++
        if n, ok := dm[na]; ok {
            res += n
        }
        dm[na]++
    }
    return res
}
