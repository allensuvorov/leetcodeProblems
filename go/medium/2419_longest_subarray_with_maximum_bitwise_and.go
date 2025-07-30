func longestSubarray(nums []int) int {
    var maxNum int

    for _, v := range nums {
        maxNum = max(maxNum, v)
    }

    var res, cnt int
    
    for _, v := range nums {
        if v == maxNum {
            cnt++
            res = max(res, cnt)
        } else {
            cnt = 0
        }
    }
    return res
}
