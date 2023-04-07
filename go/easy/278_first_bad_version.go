package easy


func firstBadVersion(n int) int {
    // ggggggbbbbbbbbb 
    l := 1
    r := n
    for l < r {
        m := (l + r)/2
        switch isBadVersion(m) {
            case true:
                r = m
            case false:
                l = m + 1
        }
    }
    return l
}
