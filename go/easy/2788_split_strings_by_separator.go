func splitWordsBySeparator(words []string, separator byte) []string {
    ans := []string{}
    word := make([]byte, 0, 20)
    for _, str := range words {
        for i := range str {
            if str[i] == separator {
                if len(word) != 0 {
                    ans = append(ans, string(word))
                    word = word[:0]
                }
            } else {
                word = append(word, str[i])
            }
        }
        if len(word) != 0 {
            ans = append(ans, string(word))
            word = word[:0]
        }

    }
    return ans
}
