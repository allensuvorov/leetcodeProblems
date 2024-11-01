func longestSubarray(nums []int) int {
    res := 0
    onesCount := 0
    zeroCount := 0
    l := 0
    r := 0

    zerosExist := false

    for r < len(nums) {
        if nums[r] == 1 {
            onesCount++
        } else {
            zeroCount++
            zerosExist = true
        }
        
        if zeroCount > 1 {
            if nums[l] == 1 {
                onesCount--
            } else {
                zeroCount--
            }
            l++
        }

        if zeroCount <= 1 {
            res = max(res, onesCount)
        }
        r++
    }
    if !zerosExist {
        res--
    }
    return res
}
