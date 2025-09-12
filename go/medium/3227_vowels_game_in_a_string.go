func doesAliceWin(s string) bool {
    vowels := "aeiou"
    vowelCount := 0
    for _, v := range s {
        if strings.ContainsRune(vowels, v) {
            vowelCount++
        }
    }
    if vowelCount == 0 {
        return false
    }
    return true
}
