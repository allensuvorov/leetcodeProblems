func findTheDifference(s string, t string) byte {
    var sum byte
    for i := range s {
        sum += t[i]
        sum -= s[i]
    }
    sum += t[len(t)-1]
    return sum
}
