func findTheDifference(s string, t string) byte {
    var sum byte
    for i := range s {
        sum += t[i]
        sum -= s[i]
    }
    sum += t[len(t)-1]
    return sum
}

// bitwise solution
func findTheDifference(s string, t string) byte {
    var addedLetter byte
    addedLetter ^= t[len(t) - 1]
    for i := range s {
        addedLetter^=s[i]
        addedLetter^=t[i]
    }
    return addedLetter
}
