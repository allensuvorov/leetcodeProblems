func timeRequiredToBuy(tickets []int, k int) int {
    time := 0
    for person, number := range tickets {
        if person > k {
            if number < tickets[k] {
                time += number
            } else {
                time += tickets[k] - 1
            }
        } else {
            if number < tickets[k] {
                time += number
            } else {
                time += tickets[k]
            }
        }
    }
    return time
}
