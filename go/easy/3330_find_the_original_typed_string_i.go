func possibleStringCount(word string) int {
    count := 1
    for i := 0; i < len(word) - 1; i++ {
        if word[i] == word[i+1] { // duplicate
            count++
        }
    }
    return count
}
