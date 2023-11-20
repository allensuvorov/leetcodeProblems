func halvesAreAlike(s string) bool {
    size := len(s)
    vSet := map[byte]bool{}
    vowels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
    for _, v := range vowels {
        vSet[v] = true
    }
    c1, c2 := 0, 0
    for i := 0; i < size/2; i ++ {
        if vSet[s[i]] {
            c1++
        }
        if vSet[s[i+size/2]] {
            c2++
        }
    }
    return c1 == c2
}
