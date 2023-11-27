func maxProductDifference(nums []int) int {
    min1, min2 := math.MaxInt, math.MaxInt
    max1, max2 := math.MinInt, math.MinInt
    
    for _, v := range nums {
        if max1 < v {
            max2, max1 = max1, v
        } else if max2 < v {
            max2 = v
        }
        if min1 > v {
            min2, min1 = min1, v
        } else if min2 > v {
            min2 = v
        }
    }
    fmt.Println(max1,max2,min1,min2)
    return max1*max2 - min1*min2
}
