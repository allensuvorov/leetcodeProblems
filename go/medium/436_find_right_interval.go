func findRightInterval(intervals [][]int) []int {
    n := len(intervals)
    starts := make([][]int, n) // pairs: start and initial index
    for i, v := range intervals {
        starts[i] = []int{v[0], i}
    }

    slices.SortFunc(starts, func(a, b []int) int {
        return a[0] - b[0]
    })

    res := make([]int, n)
    for i, v := range intervals {
        res[i] = binSearch(starts, v[1])
    }

    return res
}

func binSearch(starts [][]int, end int) int {
    n := len(starts)
    l, r := 0, n
    res := -1

    for l < r {
        m := l + (r-l)/2
        if starts[m][0] == end {
            return starts[m][1]
        }

        if starts[m][0] < end {
            l = m + 1
        }

        if starts[m][0] > end {
            r = m
            res = starts[m][1]
        }
    }
    return res
}
