func longestSubarray(nums []int) int {
    res := 0
    onesCount := 0
    zerosCount := 0
    l := 0
    r := 0

    for r < len(nums) {
        if nums[r] == 1 {
            onesCount++
        } else {
            zerosCount++
        }
        
        if zerosCount > 1 {
            if nums[l] == 1 {
                onesCount--
            } else {
                zerosCount--
            }
            l++
        }

        if zerosCount <= 1 {
            res = max(res, onesCount)
        }
        r++
    }
    if res == len(nums) {
        res--
    }
    return res
}
