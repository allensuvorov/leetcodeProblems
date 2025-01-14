func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
    aliceSum := 0
    for _, v := range aliceSizes {
        aliceSum += v
    }
    
    bobSum := 0
    bobSet := map[int]bool{}
    for _, v := range bobSizes {
        bobSum += v
        bobSet[v] = true
    }

    diff := bobSum - aliceSum

    for _, v := range aliceSizes {
        if bobSet[v + diff/2] {
            return []int{v, v + diff/2}
        }
    }
    return nil
}
