func scoreBalance(s string) bool {
    /*
    two passes solution
    */
    const offset = 'a' - 1
    var totalScore rune // rune type is an alias for int32
    for _, v := range s {
        totalScore += v - offset
    }
    var prefixScore rune
    for _, v := range s {
        prefixScore += v - offset
        if prefixScore * 2 == totalScore {
            return true
        }
    }
    return false
}
