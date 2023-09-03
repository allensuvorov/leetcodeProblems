func longestConsecutive(nums []int) int {
    res := 0
    m := map[int]bool{}
    for _, v := range nums {
        m[v] = false
    }
    
    for _, v := range nums {
        if !m[v] {
            count := 1
            num := v + 1
            ok := false
            for _, ok = m[num]; ok; _, ok = m[num] {
                m[num] = true
                count++
                num++
            }
            num = v - 1
            for _, ok = m[num]; ok; _, ok = m[num] {
                m[num] = true
                count++
                num--
            }
            if res < count {
                res = count
            }
        }
    }
    return res
}
