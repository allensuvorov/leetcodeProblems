func canConstruct(ransomNote string, magazine string) bool {
    hm := make(map[rune]int)
    for _, c := range magazine {
        hm[c]++
    }
    for _, c := range ransomNote {
        hm[c]--
        if hm[c] < 0 {
            return false
        }
    }
    return true
}
