func thirdMax(nums []int) int {
    maxNum1 := math.MinInt
    maxNum2 := math.MinInt
    maxNum3 := math.MinInt

    for _, v := range nums {
        if v != maxNum1 && v != maxNum2 && v != maxNum3 {
            if v > maxNum1 {
                maxNum3 = maxNum2
                maxNum2 = maxNum1
                maxNum1 = v
            } else if v > maxNum2 {
                maxNum3 = maxNum2
                maxNum2 = v
            } else if v > maxNum3 {
                maxNum3 = v
            }
        }
    }

    if maxNum3 > math.MinInt {
        return maxNum3 
    }
    return maxNum1
}
