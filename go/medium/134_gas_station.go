func canCompleteCircuit(gas []int, cost []int) int {
    res, tank, total := 0, 0, 0

    for i := range len(gas) {
        diff := gas[i] - cost[i]
        tank += diff
        total += diff
        if tank < 0 {
            tank = 0
            res = i + 1
        }
    }

    if total < 0 {
        return -1
    }

    return res
}
