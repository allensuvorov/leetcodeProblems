// built-in methods solution
func reverseWords(s string) string {
    words := strings.Fields(s)
    slices.Reverse(words)
    return strings.Join(words, " ")
}

// prev solution
func reverseWords(s string) string {
    res := []byte{}
    words := strings.Fields(s)
    for _, w := range words {
        if len(res) == 0 {
            res = append(res, []byte(w)...)    
        } else {
            w = w + " "
            res = append([]byte(w), res...)
        }
    }
    return string(res)
}
