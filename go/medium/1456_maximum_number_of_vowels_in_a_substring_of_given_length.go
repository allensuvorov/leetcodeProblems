func maxVowels(s string, k int) int {
    vowels := map[byte]bool{'a':true, 'e':true, 'i':true, 'o':true, 'u':true}
    maxCount := 0
    count := 0
    for i := range s {
        if vowels[s[i]] {
            count++
        }
        if i >= k && vowels[s[i-k]] {
            count--
        }
        maxCount = max(maxCount, count)
    }
    return maxCount   
}
