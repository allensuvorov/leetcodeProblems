func longestSubarray(nums []int) int {
    result := 0
    onesCount := 0
    zerosCount := 0
    
    for left, right := 0, 0; right < len(nums); right++ {
        if nums[right] == 1 {
            onesCount++
        } else {
            zerosCount++
        }
        
        if zerosCount > 1 {
            if nums[left] == 1 {
                onesCount--
            } else {
                zerosCount--
            }
            left++
        }

        if zerosCount <= 1 {
            result = max(result, onesCount)
        }
    }

    if result == len(nums) {
        result--
    }
    return result
}
