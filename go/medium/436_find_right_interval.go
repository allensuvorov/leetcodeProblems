func findRightInterval(intervals [][]int) []int {
    n := len(intervals)
    for i := range intervals {
        intervals[i] = append(intervals[i], i)
    }

    slices.SortFunc(intervals, func(a, b []int)int {
        return a[0] - b[0]
    })

    res := make([]int, n)
    for i, v := range intervals {
        res[v[2]] = binSearch(intervals[i:], v[1])
    }
    return res
}

func binSearch(intervals[][]int, end int) int {
    l, r := 0, len(intervals) - 1
    res := -1
    for l <= r {
        m := l + (r - l)/2
        if intervals[m][0] == end {
            return intervals[m][2]
        } else if intervals[m][0] > end {
            res = intervals[m][2]
            r = m - 1
        } else if intervals[m][0] < end {
            l = m + 1
        }
    }
    return res
}
