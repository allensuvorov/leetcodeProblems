func reverseVowels(s string) string {
    vowels := []byte("aeiouAEIOU")
    res := []byte(s)
    l, r := 0, len(s)-1
    for l < r {
        if !slices.Contains(vowels, res[l]) {
            l++
            continue
        }
        if !slices.Contains(vowels, res[r]) {
            r--
            continue
        }
        res[l], res[r] = res[r], res[l]
        l++
        r--
    }
    return string(res)
}
