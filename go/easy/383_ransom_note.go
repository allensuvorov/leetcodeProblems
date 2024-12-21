func canConstruct(ransomNote string, magazine string) bool {
    have := make(map[byte]int)
    for i := range magazine {
        have[magazine[i]]++
    }

    need := make(map[byte]int)
    for i := range ransomNote {
        need[ransomNote[i]]++
        if need[ransomNote[i]] > have[ransomNote[i]] {
            return false
        }
    }
    return true
}
