func executeInstructions(n int, startPos []int, s string) []int {
    // simulation O(m^2)
    ans := make([]int, len(s))
    for i := range s {
        r := startPos[0]
        c := startPos[1]
        j := 0
        for i+j < len(s) {
            switch s[i+j] {
            case 'U': r--
            case 'D': r++
            case 'L': c--
            case 'R': c++
            }
            if r >= 0 && r < n && c >= 0 && c < n {
                j++
            } else {
                break
            }
        }
        ans[i] = j
    }
    return ans
}
