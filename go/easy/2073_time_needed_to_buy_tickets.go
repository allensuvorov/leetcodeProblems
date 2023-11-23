func timeRequiredToBuy(tickets []int, k int) int {
    time := 0
    for tickets[k] > 0 { 
        for i, v := range tickets {
            if v > 0{
                tickets[i]--
                time++
            }
            if i == k && tickets[k]==0 {
                return time
            }
        }
    }
    return time
}
