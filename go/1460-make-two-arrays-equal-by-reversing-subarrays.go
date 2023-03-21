func canBeEqual(target []int, arr []int) bool {
    nums := make(map[int]int)
    for i := range target {
        nums[target[i]]++
        nums[arr[i]]--
        if nums[target[i]] == 0 {
            delete(nums, target[i])
        }
        if nums[arr[i]] == 0 {
            delete(nums, arr[i])
        }
    }
    return len(nums) == 0
}
