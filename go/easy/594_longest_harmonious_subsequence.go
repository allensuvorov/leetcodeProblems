func findLHS(nums []int) int {
    // build frequency map
    fm := make(map[int]int)
    for _, v := range nums {
        fm[v]++
    }

    res := 0
    // get max len
    for num := range fm {
        if v, ok := fm[num+1]; ok {
            res = max(res, fm[num] + v)
        }
    }
    return res
}
