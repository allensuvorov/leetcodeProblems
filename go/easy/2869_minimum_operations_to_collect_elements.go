func minOperations(nums []int, k int) int {
    numSet := map[int]struct{}{}
    var i int
    for i = len(nums) - 1; i >=0 && len(numSet) < k; i-- {
        if nums[i] <= k {
            numSet[nums[i]] = struct{}{}
        }
    }
    return len(nums) - i - 1
}
