func majorityElement(nums []int) []int {
    n := len(nums)

    fm := map[int]int{}
    for _, v := range nums {
        fm[v]++
    }

    set := map[int]bool{}
    for k, v := range fm {
        if v > n/3 {
            set[k] = true
        }
    }

    res := make([]int, 0, len(set))
    for v := range set {
        res = append(res, v)
    }
    return res
}
