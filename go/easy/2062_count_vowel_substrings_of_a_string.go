package easy

func countVowelSubstrings(word string) int {
    vm := map[byte]bool{'a':true, 'e':true, 'i':true, 'o':true, 'u':true}
    targetVS := [26]bool{('a'-'a'):true, ('e'-'a'):true, ('i'-'a'):true, ('o'-'a'):true, ('u'-'a'):true}
    
    res := 0
    for l := range word {
        currentVS := [26]bool{}
        for r := l; r < len(word) && vm[word[r]]; r++ {
            currentVS[word[r]-'a'] = true
            if currentVS == targetVS {
                res++
            }
        }
    }
    return res
}
