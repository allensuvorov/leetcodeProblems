func minSubsequence(nums []int) []int {
    res := []int{}
    // get sum
    sum := 0
    for _, v := range nums {
        sum += v
    }
    
    // find max, add to subSum, append to res, pop it from nums
    subSum := 0
    for subSum * 2 <= sum {
        maxNum := 0
        maxInd := 0
        for i, v := range nums {
            if v > maxNum {
                maxNum = v
                maxInd = i
            }
        }
        subSum += maxNum
        res = append(res, maxNum)
        nums[maxInd] = 0
    }
    return res
}
