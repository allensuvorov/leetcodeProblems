func findWords(words []string) []string {
    res := []string{}
    keyRowsStrings := [3]string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
    keyRowsMaps := [3]map[rune]bool{}
    for i := range keyRowsStrings {
        keyRowsMaps[i] = map[rune]bool{}
        for _, v := range keyRowsStrings[i] {
            keyRowsMaps[i][v] = true
            keyRowsMaps[i][v-32] = true
        }
    }
    fmt.Println(keyRowsMaps)
    
    for i := range words {
        // check the word
        flag := true
        row := 0
        for j := range keyRowsMaps {
            if _, ok := keyRowsMaps[j][rune(words[i][0])]; ok {
                row = j
            }
        }
        for _, v := range words[i] {
            if _, ok := keyRowsMaps[row][v]; !ok {
                flag = false
                break
            }
        }
        if flag {
            res = append(res, words[i])
        }
    }
    return res
}
