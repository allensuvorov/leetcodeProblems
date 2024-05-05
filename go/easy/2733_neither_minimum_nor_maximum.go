func findNonMinOrMax(nums []int) int {
    maxNum, minNum := 0, 101
    for _, v := range nums {
        maxNum = max(v, maxNum)
        minNum = min(v, minNum)
    }
    for _, v := range nums {
        if v != maxNum && v != minNum {
            return v
        }
    }
    return -1
}
