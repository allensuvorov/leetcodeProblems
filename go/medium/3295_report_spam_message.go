func reportSpam(message []string, bannedWords []string) bool {
    isBanned := map[string]bool{}
    for _, v := range bannedWords {
        isBanned[v] = true
    }

    count := 0
    for _, v := range message {
        if isBanned[v] {
            count++
        }
    }
    return count > 1
}
