func maxFreeTime(eventTime int, startTime []int, endTime []int) int {
    n := len(startTime)
    maxBreakFromL := make([]int, n)
    maxBreak := 0
    for i := range n {
        currentBreak := startTime[i]
        if i > 0 {
            currentBreak -= endTime[i-1]
        }
        maxBreak = max(maxBreak, currentBreak)
        maxBreakFromL[i] = maxBreak
    }
    
    maxBreakFromR := make([]int, n)
    maxBreak = 0
    for i := n-1; i >=0; i-- {
        currentBreak := eventTime - endTime[i]
        if i < n-1 {
            currentBreak = startTime[i+1] - endTime[i]
        }
        maxBreak = max(maxBreak, currentBreak)
        maxBreakFromR[i] = maxBreak
    }
    // fmt.Println(maxBreakFromL)
    // fmt.Println(maxBreakFromR)
    res := 0
    for i := range n {
        leftBreak := startTime[i]
        if i > 0 {
            leftBreak = startTime[i] - endTime[i-1] 
        }

        rightBreak := eventTime - endTime[i]
        if i < n - 1 {
            rightBreak = startTime[i+1] - endTime[i]
        }

        // check free space on left
        maxBreak = 0
        if i > 0 {
            maxBreak = max(maxBreak, maxBreakFromL[i-1])
        }

        if i < n - 1 {
            maxBreak = max(maxBreak, maxBreakFromR[i+1])
        }

        meetingDuration := endTime[i] - startTime[i]
        if maxBreak >= meetingDuration {
            res = max(res, leftBreak + rightBreak + meetingDuration)
        } else {
            res = max(res, leftBreak + rightBreak)
        }
    }
    return res
}
