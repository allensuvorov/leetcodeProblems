func longestSubarray(nums []int) int {
    result := 0
    zerosCount := 0
    l := 0
    for r := range nums {
        if nums[r] == 0 {
            zerosCount++
        }
        // cut tail till zeros < 2
        for zerosCount > 1 {
            if nums[l] == 0 {
                zerosCount--
            }
            l++
        }
        result = max(result, r - l)
    }
    return result
}
