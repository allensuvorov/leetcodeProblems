func substringXorQueries(s string, queries [][]int) [][]int {
    endpoints := make(map[int][]int)
    // time Big O(32*32*n) = O(n)
    for i := range s {
        num := 0
        for j := i; j < len(s) && j <= i + 30; j++ {
            bit := int(s[j] - '0')
            num = num << 1
            num = num | bit
            if _, exists := endpoints[num]; !exists {
                if num == 0 || s[i] != '0' {
                    endpoints[num] = []int{i,j}
                }
            }
        }
    }
    res := [][]int{}
    for _, v := range queries {
        target := v[0] ^ v[1]
        if v, exists := endpoints[target]; exists {
            res = append(res, v)
        } else {
            res = append(res, []int{-1,-1})
        }
    }
    return res
}
