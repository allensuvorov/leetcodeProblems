func checkIfCanBreak(s1 string, s2 string) bool {
    sl1 := []byte(s1)
    sl2 := []byte(s2)
    slices.Sort(sl1)
    slices.Sort(sl2)
    s1 = string(sl1)
    s2 = string(sl2)
    var big int = 0
    for i := range s1 {
        if s1[i] != s2[i] {
            if big == 0 {
                if s1[i] > s2[i] {
                    big = 1
                } else {
                    big = 2
                }
            } else {
                if s1[i] > s2[i] && big == 2 {
                    return false
                }
                if s1[i] < s2[i] && big == 1 {
                    return false
                }
            }
            
        }
    }
    return true
}
