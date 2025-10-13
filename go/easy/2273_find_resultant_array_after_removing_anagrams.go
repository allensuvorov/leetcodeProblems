func removeAnagrams(words []string) []string {
    prev := [26]int{} // char counts
    res := []string{}
    for _, word := range words {
        curr := [26]int{} // char counts
        for _, char := range word {
            curr[char - 'a']++
        }

        if curr != prev {
            res = append(res, word)
            prev = curr
        }
    }
    return res
}
