func insert(intervals [][]int, newInterval []int) [][]int {
    n := len(intervals)
    i := 0
    res := make([][]int, 0)
        
    for i < n && intervals[i][1] < newInterval[0] {
        res = append(res, intervals[i])
        i++
    } 
    
    for i < n && intervals[i][0] <= newInterval[1] {
        newInterval[0] = min(newInterval[0], intervals[i][0])
        newInterval[1] = max(newInterval[1], intervals[i][1])
        i++
    }

    res = append(res, newInterval)
        
    for i < n && intervals[i][0] > newInterval[1] {
        res = append(res, intervals[i])
        i++
    }
    return res
}


// 1---3   6---9
//   2---5
// 1-----5
