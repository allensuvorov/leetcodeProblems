// two mins
func increasingTriplet(nums []int) bool {
    min1 := math.MaxInt
    min2 := math.MaxInt

    for _, v := range nums {
        if v <= min1 {
            min1 = v
        } else if v <= min2 {
            min2 = v
        } else {
            return true
        }
    }
    return false
}

// min and max
func increasingTriplet(nums []int) bool {
    l, r := 0, len(nums) - 1
    minSoFar := nums[0]
    maxSoFar := nums[len(nums) - 1]
    for l < r {
        if minSoFar < nums[l] && nums[l] < maxSoFar {
            return true
        }
        if minSoFar < nums[r] && nums[r] < maxSoFar {
            return true
        }
        minSoFar = min(minSoFar, nums[l])
        maxSoFar = max(maxSoFar, nums[r])
        l++
        r--
    }
    return false
}
