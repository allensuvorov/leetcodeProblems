func frequencySort(nums []int) []int {
    const offset = 100
    freqList := make([]int, 201)
    
    for _, v := range nums {
        freqList[v + offset]++
    }

    freqTable := make([][]int, 101)
    for num, freq := range freqList {
        if freq > 0 {            
            freqTable[freq] = append(freqTable[freq], num - offset)
        }
    }

    ans := []int{}
    for freq, nums := range freqTable {
        for i := len(nums) - 1; i >= 0; i-- {
            for range freq {
                ans = append(ans, nums[i])
            }
        }
    }
    return ans
}

// ind  1, 2, 3
// freq 2, 3, 1

// 1: 2
// 2: 3
// 3: 1, 
