func fillCups(amount []int) int {
    time := 0
    slices.Sort(amount)
    for amount[2] > 0 {
        time++
        if amount[1] > 0 {
            amount[1]--
        }
        if amount[2] > 0 {
            amount[2]--
        }
        slices.Sort(amount)
    }
    return time
}
