func averageWaitingTime(customers [][]int) float64 {
    chefBusyTill := 0
    totalWaitTime := 0
    for _, v := range customers {
        if v[0] >= chefBusyTill {
            totalWaitTime += v[1]
            chefBusyTill = v[0] + v[1]
        } else {
            wait := chefBusyTill - v[0] + v[1]
            totalWaitTime += wait
            chefBusyTill = v[0] + wait
        }
    }
    return float64(totalWaitTime)/float64(len(customers))
}
