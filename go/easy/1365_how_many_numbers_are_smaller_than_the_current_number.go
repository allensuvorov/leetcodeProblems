func smallerNumbersThanCurrent(nums []int) []int {
    sortedNums := slices.Clone(nums)
    slices.Sort(sortedNums)

    uniqNumIdx := make(map[int]int) // map of unique numbers
    
    for i, v := range sortedNums {
        if _, exists := uniqNumIdx[v]; !exists {
            uniqNumIdx[v] = i
        }
    }

    res := make([]int, len(nums))
    for i, v := range nums {
        res[i] = uniqNumIdx[v]
    }
    return res
}
