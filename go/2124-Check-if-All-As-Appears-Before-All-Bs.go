func checkString(s string) bool {
    // look for a after first b
    for i := 0; i < len(s); i++ {
        if s[i] == 'b' {
            for ; i < len(s); i++ {
                if s[i] == 'a' {
                    return false
                }
            }
        }
    }
    return true
}
