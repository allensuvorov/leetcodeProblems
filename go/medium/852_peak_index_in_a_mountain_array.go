package medium

func peakIndexInMountainArray(arr []int) int {
    l := 0
    r := len(arr)-1
    for l <= r {
        m := (l + r) >> 1
        switch {
        case m == 0: l++
        case m == len(arr)-1: r--
        case arr[m] > arr[m-1] && arr[m] > arr[m+1]: return m
        case arr[m] > arr [m-1]: l = m + 1
        case arr[m] < arr [m-1]: r = m - 1
        }
    }
    return l
}
