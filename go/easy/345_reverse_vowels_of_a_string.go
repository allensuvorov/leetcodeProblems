func reverseVowels(s string) string {
    a := []byte(s)
    vowels := map[byte]bool{'a':true, 'e':true, 'i':true, 'o':true, 'u':true, 'A':true, 'E':true, 'I':true, 'O':true, 'U':true}
    for l, r := 0, len(a) - 1; l < r; {
        if !vowels[a[l]] {
            l++
            continue
        }
        if !vowels[a[r]] {
            r--
            continue
        }
        a[l], a[r] = a[r], a[l]
        l++
        r--
    }
    return string(a)
}
