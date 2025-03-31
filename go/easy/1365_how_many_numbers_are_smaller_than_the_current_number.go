func smallerNumbersThanCurrent(nums []int) []int {
    freqMap := make([]int, 101)
    for _, v := range nums {
        freqMap[v]++
    }
    
    smallerNumCount := make([]int, 101)
    for i := range freqMap{
        if i > 0 {
            smallerNumCount[i] = freqMap[i-1] + smallerNumCount[i-1]
        }
    }

    res := make([]int, len(nums)) 
    
    for i, v := range nums {
        res[i] = smallerNumCount[v]
    }

    return res
}

