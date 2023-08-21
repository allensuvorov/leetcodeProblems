func canConstruct(ransomNote string, magazine string) bool {
    cf := [26]rune{}
    for _, v := range magazine {
        cf[v-'a']++
    }
    for _, v := range ransomNote {
        cf[v-'a']--
        if cf[v-'a'] < 0 {
            return false
        }
    }
    return true
}
