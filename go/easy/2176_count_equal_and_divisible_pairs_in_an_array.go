func countPairs(nums []int, k int) int {
    valueToIndices := make(map[int][]int) //
    res := 0
    for i, v := range nums {
        for _, index := range valueToIndices[v] {
            if index * i % k == 0 {
                res++
            }
        }
        valueToIndices[v] = append(valueToIndices[v], i)
    }
    return res
}
