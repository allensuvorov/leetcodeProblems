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
