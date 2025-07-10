func maxFreeTime(eventTime int, startTime []int, endTime []int) int {
    n := len(startTime)
    res := 0

    getLR := func(i int) (int, int) {
        l := 0
        if i > 0 {
            l = endTime[i-1]
        }
        r := eventTime
        if i < n - 1 {
            r = startTime[i+1]
        }
        return l, r
    }

    maxSlotSoFar := 0
    for i := range n {
        l, r := getLR(i)
        meeting := endTime[i] - startTime[i]
        if maxSlotSoFar >= meeting {
            res = max(res, r - l)
        } else {
            res = max(res, r - l - meeting)
        }
        maxSlotSoFar = max(maxSlotSoFar, startTime[i] - l)
    }
    
    maxSlotSoFar = 0
    for i := n-1; i >=0; i-- {
        l, r := getLR(i)
        meeting := endTime[i] - startTime[i]
        if maxSlotSoFar >= meeting {
            res = max(res, r - l)
        } else {
            res = max(res, r - l - meeting)
        }
        maxSlotSoFar = max(maxSlotSoFar, r - endTime[i])
    }

    return res
}
