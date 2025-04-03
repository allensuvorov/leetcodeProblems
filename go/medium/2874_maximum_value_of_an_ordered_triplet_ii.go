func maximumTripletValue(nums []int) int64 {
	// 12,6,1,2,7

    prefMax := make([]int, len(nums)) // 12,12,12,12,12
    suffMax := make([]int, len(nums)) // 12,7,7,7,7

    maxVal := 0
    for i, v := range nums {
        if v > maxVal {
            prefMax[i] = v
            maxVal = v
        } else {
            prefMax[i] = maxVal
        }
    }

    maxVal = 0
    for i := len(nums)-1; i >=0; i-- {
        if nums[i] > maxVal {
            suffMax[i] = nums[i]
            maxVal = nums[i]
        } else {
            suffMax[i] = maxVal
        }

    }

    var res int64
	  var curVal int64

    for j := 1; j < len(nums) - 1; j++ {
        i := j - 1
        k := j + 1
        curVal = int64(prefMax[i] - nums[j]) * int64(suffMax[k])
        res = max(res, curVal)
    }
	
	return maxVal
}
