func countPrefixSuffixPairs(words []string) int64 {
    var count int64 = 0
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

// HARD LEVEL
// n*n is 10^10 thats too slow

// Optimization ideas:
//   prefsuff can not be longer than the target
//   source is of the left side of the target
//   some kind of sorting

// ["a","aba","ababa","aa"]
