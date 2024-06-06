func dominantIndex(nums []int) int {
    ans := 0
    maxNum := 0
    secMaxNum := 0
    for i, v := range nums {
        if v > secMaxNum {
            secMaxNum = v
            if v > maxNum {
                secMaxNum = maxNum
                maxNum = v
                ans = i
            } 
        }
    }
    if maxNum >= secMaxNum * 2 {
        return ans
    }
    return - 1
}
