func fullJustify(words []string, maxWidth int) []string {
    lineLength := 0
    i := 0
    for i < len(words) {
        if lineLength + len(words[i]) + i <= maxWidth {
            lineLength += len(words[i])
        } else {
            break
        }
        i++
    }
    
    // base case
    if i == len(words) {
        return []string{leftJustify(words[:i], maxWidth)}
    }
    
    currentLine := pad(words[:i], maxWidth)
    remainingLines := fullJustify(words[i:], maxWidth)
    return append([]string{currentLine}, remainingLines...)
}

func leftJustify(words []string, maxWidth int) string {
    line := words[0]
    for i := 1; i < len(words); i++ {
        line += " " + words[i]
    }
    for range maxWidth - len(line) {
        line += " "
    }
    return line
}

func pad(words []string, maxWidth int) string {
    line := ""
    if len(words) == 1 {
        line = words[0]
        for range maxWidth - len(line) {
            line += " "
        }
        return line
    }

    wordsLen := 0
    for _, v := range words {
        wordsLen += len(v)
    }

    padLen := maxWidth - wordsLen
    padCount := len(words) - 1

    minPadsLen := padLen / padCount
    extraPads := padLen % padCount


    for i := 0; i < len(words) - 1; i++ {
        for range minPadsLen {
            words[i] += " "
        }
        if i < extraPads {
            words[i] += " "
        }
    }

    for _, v := range words {
        line += v
    }
    return line
}
