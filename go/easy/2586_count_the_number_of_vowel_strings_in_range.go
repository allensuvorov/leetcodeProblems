func vowelStrings(words []string, left int, right int) int {
    count := 0
    vowels := map[byte]bool{'a':true, 'e': true, 'i': true, 'o': true, 'u': true}
    for i := left; i <= right; i++ {
        if vowels[words[i][0]] && vowels[words[i][len(words[i])-1]] {
            count++
        }
    }
    return count
}
