func minimumPairRemoval(nums []int) int {
    count := 0
    for hasDecrease(nums) {
        doOperation(nums)
        nums = nums[:len(nums)-1]
        count++
    }
    return count
}

func hasDecrease(nums []int) bool {
    for i := 0; i + 1 < len(nums); i++ {
        if nums[i] > nums[i+1] {
            return true
        }
    }
    return false
}

func doOperation(nums []int) {
    // find min sum of a pair
    minSum := nums[0] + nums[1]
    for i := 0; i + 1 < len(nums); i++ {
        minSum = min(minSum, nums[i]+nums[i+1])
    }
    // find leftmost pair with the min sum and replace
    for i := 0; i + 1 < len(nums); i++ {
        if nums[i] + nums[i+1] == minSum {
            replacePair(nums, i)
            return
        }
    }
}

func replacePair(nums []int, i int) {
    nums[i] = nums[i] + nums[i+1]
    for i++ ; i + 1 < len(nums) ; i++ {
        nums[i] = nums[i+1]
    }
}
