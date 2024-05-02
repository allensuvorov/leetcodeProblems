func reversePrefix(word string, ch byte) string {
    pref := make([]byte, len(word))
    for l := range word {
        r := len(pref) - 1 - l
        pref[r] = word[l]
        if word[l] == ch {
            return string(pref[r:]) + word[l+1:]
        }
    }
    return word
}
