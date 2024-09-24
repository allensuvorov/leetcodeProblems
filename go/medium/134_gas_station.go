func canCompleteCircuit(gas []int, cost []int) int {
    // bruteforce, simulation: 10^5; n^2 = 10^10
    // answer is right after the lowest deficit point, if we ever get out of deficit
    worstDeficitPoint := 0
    worstDeficit := 10_000
    totalDeficit := 0
    for i := range gas {
        totalDeficit += gas[i] - cost[i] 
        fmt.Println(totalDeficit)
        if totalDeficit <= worstDeficit {
            worstDeficit = totalDeficit
            worstDeficitPoint = i
        }
    }

    if totalDeficit < 0 {
        return -1
    }

    if worstDeficitPoint == len(gas) - 1 {
        return 0
    }

    return worstDeficitPoint + 1
}
