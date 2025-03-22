func substringXorQueries(s string, queries [][]int) [][]int {
    endpoints := make(map[string][]int)
    // time Big O(32*32*n) = O(n)
    for i := range s {
        for j := i+1; j <= len(s) && j <= i + 30; j++ {
            subStr := s[i:j]
            if _, exists := endpoints[subStr]; !exists {
                endpoints[subStr] = []int{i,j-1}
            }
        }
    }
    res := [][]int{}
    for _, v := range queries {
        target := v[0] ^ v[1]
        if v, exists := endpoints[strconv.FormatInt(int64(target), 2)]; exists {
            res = append(res, v)
        } else {
            res = append(res, []int{-1,-1})
        }
    }
    return res
}
