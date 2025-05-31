// for one string s
func isSubsequence(s string, t string) bool {
    need := 0
    for have := range t {
        if need < len(s) && t[have] == s[need] {
            need++
        }
    }
    return need == len(s)
}
