func repeatedSubstringPattern(s string) bool {
    return strings.Contains((s+s)[1:len(s+s)-1], s)
}
