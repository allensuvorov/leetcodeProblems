func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
    repeatingNums := []int{}

    set1 := freqCount(nums1)
    set2 := freqCount(nums2)
    set3 := freqCount(nums3)
    
    for num := 1; num <= 100; num++ {
        if set1[num] + set2[num] + set3[num] > 1 {
            repeatingNums = append(repeatingNums, num)
        }
    }
    return repeatingNums
}

func freqCount(nums []int) []int {
    freqCount := make([]int, 101)
    for _, v := range nums {
        if freqCount[v] == 0 {
            freqCount[v] = 1
        }
    }
    return freqCount
}
