func findLucky(arr []int) int {
    m := map[int]int{}
    max := -1
    for _, v := range arr {
        m[v]++
    }
    for k, v := range m {
        if k == v && max < k {
            max = k
        }
    }
    return max
}
