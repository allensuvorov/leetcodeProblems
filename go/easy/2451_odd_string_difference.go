func oddString(words []string) string {
    same := ""
    diffWord := make(map[[20]byte]string, 2) // map diff with word
    for _, word := range words {
        curDiff := [20]byte{}
        for j := 0; j < len(word) - 1; j++ {
            curDiff[j] = word[j+1]-word[j]
        }
        if diffWord[curDiff] != "" {
            same = word 
        }
        diffWord[curDiff] = word
    }

    for _, word := range diffWord {
        if word != same {
            return word
        }
    }
    return ""
}
