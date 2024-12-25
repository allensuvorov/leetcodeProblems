func reverseVowels(s string) string {
    res := []byte(s)
    vowels := map[byte]bool {
        'a':true, 'e':true, 'i':true, 'o':true, 'u':true,
        'A':true, 'E':true, 'I':true, 'O':true, 'U':true,
    }
    for l, r := 0, len(res)-1; l<r; {
        if !vowels[res[l]] {
            l++
            continue
        }
        if !vowels[res[r]] {
            r--
            continue
        }
        res[l], res[r] = res[r], res[l]
        l++
        r--
    }
    return string(res)
}
