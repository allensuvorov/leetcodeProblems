func numWaterBottles(numBottles int, numExchange int) int {
    ans := 0
    emptyBottles := 0
    for numBottles > 0 {
        ans += numBottles
        emptyBottles += numBottles
        numBottles = emptyBottles / numExchange
        emptyBottles = emptyBottles % numExchange
    }
    return ans
}
