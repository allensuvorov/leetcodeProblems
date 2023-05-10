func kthDistinct(arr []string, k int) string {
    m := map[string]int{}
    for i := range arr {
        m[arr[i]]++
    }
    for i := range arr {
        if m[arr[i]] == 1 {
            k--
        }
        if k == 0 {
            return arr[i]
        }
    }
    return ""
}
