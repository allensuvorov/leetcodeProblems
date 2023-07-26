func romanToInt(s string) int {
    m := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
    ans := 0
    l := len(s)
    for i := 0; i < l - 1; i++{
        r := s[i]
        if m[r] < m[s[i+1]] {
            ans -= m[r]
        } else {
            ans += m[r]
        }
    } 
    ans += m[s[l-1]]
    return ans
}
