func targetIndices(nums []int, target int) []int {
    sort.SliceStable(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })
    i := sort.Search(len(nums), func(i int) bool {
        return nums[i] >= target
    })
    result := []int{}
    for i < len(nums) && nums[i] == target {
        result = append(result, i)
        i ++
    }
    return result
}
