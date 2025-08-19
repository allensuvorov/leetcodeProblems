func maxVowels(s string, k int) int {
    maxVowelCount := 0
    windowVowelCount := 0
    vowels := []byte{'a', 'e', 'i', 'o', 'u'}
    for i := range s {
        if slices.Contains(vowels, s[i]) {
            windowVowelCount++
        }

        if i >= k && slices.Contains(vowels, s[i - k]) { // we have window and a vowel tail
            windowVowelCount-- // update the count
        }
        
        if i + 1 >= k { // we have a window
            maxVowelCount = max(maxVowelCount, windowVowelCount)
        }
    }
    return maxVowelCount
}
