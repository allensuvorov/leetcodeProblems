func wordPattern(pattern string, s string) bool {
    hm := make(map[rune]string)
    p := 0
    for _, v := range pattern {
        w := []byte{}
        for i := p; i < len(s); i++{
            if s[i] == ' '{
                break
            } else {
                w = append(w, s[i])
            }
        }
        p += len(w)+1
        word := string(w)
        fmt.Println(v, word)
        if _, ok := hm[v]; ok {
            if hm[v] != word {
                return false
            }
        } else {
            for _, v := range hm {
                if v == word {
                    return false
                }
            }
            hm[v] = word
        }

    }
    if p-1 != len(s) { 
        return false
    }
    return true
}
