func largestAltitude(gain []int) int {
    maxAlt := 0
    curAlt := 0
    for i := 0; i < len(gain); i++ {
        curAlt += gain[i]
        maxAlt = max(maxAlt, curAlt)
    }
    return maxAlt
}
