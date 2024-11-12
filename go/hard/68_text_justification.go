func fullJustify(words []string, maxWidth int) []string {
    lineLength := 0
    i := 0
    for i < len(words) {
        if lineLength + len(words[i]) + i < maxWidth {
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
    return append(currentLine, remainingLines...)
}

func leftJustify(words []string, maxWidth int) string {

}

func pad(words []string, maxWidth int) string {

}
