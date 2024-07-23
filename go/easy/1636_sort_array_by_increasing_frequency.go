func frequencySort(nums []int) []int {
    const offset = 100
  
    frequencyCountList := make([]int, 201)
    for _, v := range nums {
        frequencyCountList[v + offset]++
    }

    frequencyTable := make([][]int, 101)
    for num, count := range frequencyCountList {
        if count > 0 {
            if frequencyTable[count] == nil {
                nums := make([]int, 201)
                nums[num] = 1 
                frequencyTable[count] = nums
            } else {
                frequencyTable[count][num] = 1
            }
        }
    }
  
    ans := []int{}
    for frequency, nums := range frequencyTable {
        for num := len(nums) - 1; num >= 0; num-- {
            if nums[num] == 1 {
                for range frequency {
                    ans = append(ans, num - offset)
                }
            }
        }
    }
    return ans
}
