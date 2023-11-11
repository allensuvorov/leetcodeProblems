func stringMatching(words []string) []string {
    // O (n^2)
    res := []string{}
    for i := range words {
        for j := range words {
            if i != j && strings.Contains(words[j], words[i]){
                res = append(res, words[i])
                break
            }
        }
    }
    return res
}
