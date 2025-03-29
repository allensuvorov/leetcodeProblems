func getSneakyNumbers(nums []int) []int {
    counts := make([]int, len(nums)-2)
    res := []int{}
    for _, v := range nums {
        counts[v]++
        if counts[v] == 2 {
            res = append(res, v)
        }
    }
    return res
}
