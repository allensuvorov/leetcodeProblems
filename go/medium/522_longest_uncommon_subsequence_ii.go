 // solution - brut 
func findLUSlength(strs []string) int {
    max := -1
    for i := 0; i < len(strs); i++ {
        l := len(strs[i])
        for j := 0; j<len(strs); j++{
            if i == j {
                continue
            }
            if isSubSeq (strs[i], strs[j]) {
                l = -1
                break
            }
        }
        if max < l {
            max = l
        }
    }
    return max
}

func isSubSeq (a, b string) bool {
    if a == b {
        return true
    }

    if len(a) >= len(b) {
        return false
    }

    i := 0
    j := 0
    for i < len(a) && j < len(b) {
        if a[i] == b[j] {
            i++
        }
        j++
    }
    return i == len(a)
}
