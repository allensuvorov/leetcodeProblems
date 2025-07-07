func maxEvents(events [][]int) int {
    // preprocessing
    slices.SortFunc(events, func(a, b []int) int{
        // sort by end
        return a[1] - b[1]
    })

    attendedCount := 0
    today := 1 // the day we have

    for _, event := range events {
        start, end := event[0], event[1]
        if end >= today {
            attendedCount++
            if start <= today {
                today++
            } else {
                today = start + 1
            }
        }        
    }
    return attendedCount
}


// 12345
//  --
//  --
// -----
// -----
// -----
