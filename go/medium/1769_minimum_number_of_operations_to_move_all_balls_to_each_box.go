func minOperations(boxes string) []int {
    lSum := make([]int, len(boxes))
    rSum := make([]int, len(boxes))

    ballCount := 0
    prev := 0  
    for i := 0; i < len(boxes); i++ {
        lSum[i] = lSum[prev] + ballCount * (i - prev)
        if lSum[i] > 0 || boxes[i] > 0 { // positions of last box with balls
            prev = i
        }
        ballCount += int(boxes[i]-'0')
    }

    ballCount = 0           // 1
    prev = len(boxes) - 1   // end
    for i := len(boxes) - 1; i >= 0; i-- {
        rSum[i] = rSum[prev] + ballCount * (prev - i)
        if rSum[i] > 0 || boxes[i] > 0 {
            prev = i
        }
        ballCount += int(boxes[i]-'0')
    } 

    res := make([]int, len(boxes))
    for i := range res {
        res[i] = lSum[i] + rSum[i]
    }
    return res
}
