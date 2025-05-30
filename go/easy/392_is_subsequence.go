func isSubsequence(s string, t string) bool {
    j := 0
    for i := 0; i < len(t) && j < len(s); i++ { // this is what we have, one by one
        if t[i] == s[j] { // is it required?
            j++ // good, check next
        }
    }
    return j == len(s) // did we find all required?
}


func isSubsequence(s string, t string) bool {
    for i, j := 0, 0; i < len(s); i++{
        for j < len(t) && t[j] != s[i] {
            j++
        }
        if j == len(t) {
            return false
        }
        j++
    }
    return true
}
