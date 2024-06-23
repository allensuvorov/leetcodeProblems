func diStringMatch(s string) []int {
    ans := make([]int, len(s) + 1)
    l, r := 0, len(ans) - 1
    for i := range s {
        if s[i] == 'I' {
            ans[i] = l
            l++
        } else {
            ans[i] = r
            r--
        }
        if i == len(s) - 1 {
            if s[i] == 'I' {
                ans[i+1] = r
            } else {
                ans[i+1] = l
            }
        }
    }
    return ans
}

// 0, 1, 2, 3, 4
// l           
//             r
