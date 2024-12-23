func insert(intervals [][]int, newInterval []int) [][]int {
    n := len(intervals) 
    res := [][]int{}
    i := 0

    // Case 1: No overlappint before merging intervals
    for i < n && intervals[i][1] < newInterval[0]{
        res = append(res, intervals[i])
        i++
    }

    // Case 2: Overlapping and merging intervals

    for i < n && intervals[i][0] <= newInterval[1] {
        newInterval[0] = min(intervals[i][0], newInterval[0]) 
        newInterval[1] = max(intervals[i][1], newInterval[1])
        i++
    }

    res = append(res, newInterval)

    // Case 3: No overlappint after merging intervals
    for i < n {
        res = append(res, intervals[i])
        i++
    }
    
    return res
}
