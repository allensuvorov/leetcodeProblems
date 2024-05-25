func countPrefixSuffixPairs(words []string) int {
    // O(n*n) solution
    count := 0
    for i := range words {
        for j := i + 1; j < len(words); j++ {
            if isPrefixAndSuffix(words[i], words[j]) {
                count++
            }
        }
    }
    return count
}

func isPrefixAndSuffix(str1, str2 string) bool {
    if len(str1) <= len(str2) &&
        str1 == str2[:len(str1)] && 
        str1 == str2[len(str2)-len(str1):] {
        return true
    }
    return false
}
