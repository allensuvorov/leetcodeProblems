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

// 
func maxVowels(s string, k int) int {
    vowels := map[byte]bool{'a':true, 'e':true, 'i':true, 'o':true, 'u':true}
    res := 0
    count := 0
    for i := range s {
        if vowels[s[i]] {
            count++
        }

        if i >= k - 1 {
            res = max(res, count)
            if vowels[s[i-(k-1)]] {
                count--
            }
        }
    }
    return res
}
