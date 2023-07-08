func largestAltitude(gain []int) int {
    max := 0
    sum := 0
    for i := range gain {
        sum += gain[i]
        if max < sum {
            max = sum
        }
    }
    return max
}
