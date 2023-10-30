func findSpecialInteger(arr []int) int {
    n := len(arr)
    q := n/4
    for i := 0; i < n - q; i ++ {
        if arr[i] == arr[i+q] {
            return arr[i]
        }
    }
    return -1
}
