func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
    repeatingNums := []int{}

    set1 := getSet(nums1)
    set2 := getSet(nums2)
    set3 := getSet(nums3)
    
    for num := 1; num <= 100; num++ {
        if set1[num] + set2[num] + set3[num] > 1 {
            repeatingNums = append(repeatingNums, num)
        }
    }
    return repeatingNums
}

func getSet(nums []int) []int {
    set := make([]int, 101)
    for _, v := range nums {
        if set[v] == 0 {
            set[v] = 1
        }
    }
    return set
}
