func maxFreqSum(s string) int {
    counts := make(map[rune]int)
    vowels := map[rune]bool{'a':true, 'e':true, 'i':true, 'o':true, 'u':true}
    for _, char := range s {
        counts[char]++
    }

    maxVF := 0
    maxCF := 0

    for char, count := range counts {
        if vowels[char] {
            maxVF = max(maxVF, count)
        } else {
            maxCF = max(maxCF, count)
        }
    }    
    return maxVF + maxCF
}
