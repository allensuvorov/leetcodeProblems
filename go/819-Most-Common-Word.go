func mostCommonWord(paragraph string, banned []string) string {
    
    hmNonBanned := map[string]int{}
    hmBanned := map[string]bool{}
    for i := range banned {
        hmBanned[banned[i]] = true
    }
    fmt.Println("hmBanned", hmBanned)

    seps := map[rune]bool{}
    for _, v := range "!?',;. " {
        seps[v] = true
    }
    fmt.Println("seps", seps)

    word := []rune{}
    for i, v := range paragraph {
        if _, ok := seps[v]; ok {
            if len(word) != 0 {
                addIfWordNotBanned(word, hmNonBanned, hmBanned)
                word = []rune{}
            }
        } else {
            word = append(word, v)
            if i == len(paragraph)-1 {
                addIfWordNotBanned(word, hmNonBanned, hmBanned)
            }
        }
    }

    fmt.Println(hmNonBanned)
    max := 0
    var res string
    for k, v := range hmNonBanned{
        if v > max {
            max = v
            res = k
        }
    }
    return res
}

func addIfWordNotBanned (word []rune, hmNonBanned map[string]int, hmBanned map[string]bool) {
    cleanWord := strings.ToLower(string(word))
    _, ok := (hmBanned)[cleanWord]
    if !ok {
        (hmNonBanned)[cleanWord]++
        fmt.Println(cleanWord)
    }
}
