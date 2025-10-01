func numWaterBottles(numBottles int, numExchange int) int {
    res := 0
    emptyBottles := 0
    for numBottles > 0 {
        res += numBottles // drink
        emptyBottles += numBottles // they are all empty now
        numBottles = emptyBottles / numExchange // exchange
        emptyBottles = emptyBottles % numExchange // empty
    }
    return res
}
