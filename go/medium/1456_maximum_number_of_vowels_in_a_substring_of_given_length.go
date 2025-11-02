func maxVowels(s string, k int) int {
    res := 0
    curVowels := 0
    vowels := "aeiou"

    isVowel := make(map[byte]bool)
    for i := range vowels{
        isVowel[vowels[i]] = true
    }
    
    for i := range s {
        // add head item
        if isVowel[s[i]] {
            curVowels++
        }
        // cut tail item
        if i >= k && isVowel[s[i - k]] { // got a tail letter and it is a vowel
            curVowels--
        }
        
        if i >= k - 1 { // we have a window
            res = max(res, curVowels)
        }
    }
    return res
}
