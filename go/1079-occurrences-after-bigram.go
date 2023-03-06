func findOcurrences(text string, first string, second string) []string {
    words := strings.Split(text, " ")
    res := []string{}

    for i := range words {
        if words[i] == first && i+1 < len(words) {
            if words[i+1] == second && i+2 < len(words) {
                res = append(res, words[i+2])
            }
        }
    }
    return res
}
