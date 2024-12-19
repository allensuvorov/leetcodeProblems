// built-in methods solution
func reverseWords(s string) string {
    words := strings.Fields(s)
    slices.Reverse(words)
    return strings.Join(words, " ")
}

// no built-in methods solution
func reverseWords(s string) string {
    res := ""
    word := ""
    for i, v := range s {
        if v != ' ' {
            word += string(v)
        } 
        if (v == ' ' || i == len(s)-1) && len(word) > 0 {
            if len(res) == 0 {
                res = word
            } else {
                res = word + " " + res
            }
            word = ""
        }
    }
    return res
}
