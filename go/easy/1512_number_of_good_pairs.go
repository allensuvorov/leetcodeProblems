func numIdenticalPairs(nums []int) int {
    // hash frequency map, then arith progression sum // n*(n-1)/2
    sum := 0
    fm := map[int]int{}
    for _, v := range nums {
        fm[v]++
    }
    for _, v := range fm {
        sum += v*(v-1)/2
    }
    return sum
}
