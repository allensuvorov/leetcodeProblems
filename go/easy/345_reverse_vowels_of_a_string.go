func reverseVowels(s string) string {
    l, r := 0, len(s) - 1
    data := []byte(s)
    for l < r {
        if !isVowel(data[l]) {
            l++
            continue
        }
        if !isVowel(data[r]) {
            r--
            continue
        }

        data[l], data[r] = data[r], data[l]
        l++
        r--
    }
    return string(data)
}

func isVowel(c byte) bool {
    vowels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
    for _, v := range vowels {
        if c == v {
            return true
        }
    }
    return false
}
