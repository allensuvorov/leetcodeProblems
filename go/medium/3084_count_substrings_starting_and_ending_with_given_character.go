func countSubstrings(s string, c byte) int64 {
    var countChar int64 = 0
    for i := range s {
        if s[i] == c {
            countChar++
        }
    }
    var countSubstr int64 = 0
    for i := range countChar { // from 0 to countChar - 1
        countSubstr = countSubstr + 1 + i
    }
    return countSubstr
}

// a - 0 + 1 + 0 = 1
// aa - 1 + 1 + 1 = 3
// aaa - 3 + 1 + 2 = 6
// aaaa - 6 + 1 + 3 = 11
