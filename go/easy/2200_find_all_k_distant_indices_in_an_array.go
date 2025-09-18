func findKDistantIndices(nums []int, key int, k int) []int {
    n := len(nums)
    var result []int
    r := -1
    for i := range nums {
        if nums[i] == key {
            l := max(i - k, r + 1)
            r = min(n-1, i + k)
            for j := l; j <= r; j++ {
                result = append(result, j)
            }
        }
    }
    return result
}
