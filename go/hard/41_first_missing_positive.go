func firstMissingPositive(nums []int) int {
    // pre processing
    // if there a 1 in the nums?
    oneExists := false
    for _, v := range nums {
        if v == 1 {
            oneExists = true
        }
    }

    if !oneExists {
        return 1
    }

    // clean nums from negative and zeros numbers, replace them with 1s
    n := len(nums)
    for i := range nums {
        if nums[i] <= 0 || nums[i] > n {
            nums[i] = 1
        }
    }
    
    // to use the array as a hash set, no negative numbers here now
    // we will reverse to negative, those that exists
    // index 0   is 1
    // index n-1 is n, offset is 1

    // 1, 2, 1  

    for _, v := range nums {
        index := abs(v) - 1
        if nums[index] > 0 {
            nums[index] = -nums[index]
        }
    }

    for num := 1; num <= n; num++ {
        if nums[num-1] > 0 {
            return num
        }
    }

    return n + 1

}

func abs(num int) int {
    if num < 0 {
        return -num
    }

    return num
}
