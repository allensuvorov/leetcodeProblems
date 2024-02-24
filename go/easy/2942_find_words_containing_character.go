func findWordsContaining(words []string, x byte) []int {
    ans := []int{}
    for i := range words {
        for j := range words[i] {
            if words[i][j] == x {
                ans = append(ans, i)
                break
            }
        }
    }
    return ans
}
