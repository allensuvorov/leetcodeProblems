func minimumAbsDifference(arr []int) [][]int {
    var (
        n int = len(arr)
        minDiff int = math.MaxInt
        res [][]int
    )

    slices.Sort(arr)
    
    for i := 1; i < n; i++ {
        minDiff = min(minDiff, arr[i] - arr[i-1])
    }

    for i := 1; i < n; i++ {
        if arr[i] - arr[i-1] == minDiff {
            res = append(res, []int{arr[i-1], arr[i]})
        }
    }
    return res
}
