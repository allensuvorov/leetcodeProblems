func numIdenticalPairs(nums []int) int {
    // hash frequency map
    sum := 0
    fm := map[int]int{}
    for _, v := range nums {
        sum += fm[v]
        fm[v]++
    }
    return sum
}
