func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
    // find max size sliding window of K meetings
    n := len(startTime)
    result := 0
    windowFreeTime := startTime[0]

    for r := 0; r < n; r++ {
        // add leading free time
        leadingFreeTime := eventTime - endTime[r]
        if r < n - 1 {
            leadingFreeTime = startTime[r + 1] - endTime[r]
        }

        windowFreeTime += leadingFreeTime
        
        // remove trailing free time
        if l := r - k; l >= 0 {
            
            trailingFreeTime := startTime[l]
            if l > 0 {
                trailingFreeTime -= endTime[l - 1]
            }
            windowFreeTime -= trailingFreeTime
        }
        
        result = max(result, windowFreeTime)   
    }

    return result
}
