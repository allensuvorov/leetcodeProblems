func totalFruit(fruits []int) int {
    counts := map[int]int{}
    res := 0
    for l, r := 0, 0; r < len(fruits); r++ {
        counts[fruits[r]]++
        for ; len(counts) > 2; l++ {
            counts[fruits[l]]--
            if counts[fruits[l]] == 0 {
                delete(counts, fruits[l])
            }
        }
        res = max(res, r - l + 1)
    }
    return res
}
