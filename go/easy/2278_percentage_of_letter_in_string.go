func percentageLetter(s string, letter byte) int {
    size := len(s)
    count := 0
    for i := 0; i < size; i++ {
        if s[i] == letter {
            count++
        }
    }
    return count * 100/size
}
