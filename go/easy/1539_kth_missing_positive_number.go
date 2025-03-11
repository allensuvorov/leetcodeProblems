func findKthPositive(arr []int, k int) int {
    countMissing := 0
    i := 0 // index in array
    n := 1 // missingNum
    for ; countMissing < k; n ++ {

        switch {
            case i >= len(arr):
                countMissing ++
            case n != arr[i]:
                countMissing ++
            case n == arr[i]:
                i++
        }

    }
    return n - 1
}

// log n: (k+m) magic + binary search

func findKthPositive(arr []int, k int) int {
    l, r := 0, len(arr) - 1
    res := len(arr) + k
    for l <= r {
        m := l + (r-l)/2
        
        if arr[m] > k + m {
            res = k + m
            r = m - 1
        } else {
            l = m + 1
        }
    }
    return res
}
