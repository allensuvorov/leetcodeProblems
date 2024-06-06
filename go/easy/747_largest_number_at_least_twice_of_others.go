func dominantIndex(nums []int) int {
    ans := 0
    maxNum := 0
    secMaxNum := 0
    for i, v := range nums {
        if v > maxNum {
            secMaxNum = maxNum
            maxNum = v
            ans = i
        } else if v > secMaxNum {
            secMaxNum = v
        }
    }
    if maxNum >= secMaxNum * 2 {
        return ans
    }
    return - 1
}
