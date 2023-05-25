func capitalizeTitle(title string) string {
    word := []rune{}
    capTitle := make([]rune, 0, len(title))
    for _, v := range title {
        if v == ' ' {
            word = capWord(word)
            word = append(word, ' ')
            capTitle = append(capTitle, word...)
            word = []rune{}
        } else {
            word = append(word, v)
        }
    }
    word = capWord(word)
    capTitle = append(capTitle, word...)
    return string(capTitle)
}

func capWord(word []rune) []rune {
    if len(word) < 3 {
        for i, v := range word {
            word[i] = unicode.ToLower(v)
        }
        return word
    } else {
        word[0] = unicode.ToUpper(word[0])
        for i := 1; i < len(word); i ++ {
            word[i] = unicode.ToLower(word[i])
        }
        return word
    }
    return nil
}
