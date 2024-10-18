func maxVowels(s string, k int) int {
    maxCount := 0
    count := 0
    vowels := map[byte]bool{'a':true, 'e':true, 'i':true, 'o':true, 'u':true}
    for head := 0; head < len(s); head++ {
        if vowels[s[head]] {
            count++
        }
        if head >= k && vowels[s[head-k]] {
            count--
        }
        maxCount = max(maxCount, count)
    }
    return maxCount
}
