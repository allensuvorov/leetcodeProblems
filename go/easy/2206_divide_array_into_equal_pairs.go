func divideArray(nums []int) bool {
    m := make(map[int]int)
    for _, v := range nums{
        m[v]++
        if m[v] == 2 {
            delete(m, v)
        }
    }
    return len(m) == 0
}
