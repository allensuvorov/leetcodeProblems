func largestAltitude(gain []int) int {
    maxAlt := 0
    curAlt := 0
    for _, v := range gain {
        curAlt += v
        maxAlt = max(maxAlt, curAlt)
    }
    return maxAlt
}
