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
