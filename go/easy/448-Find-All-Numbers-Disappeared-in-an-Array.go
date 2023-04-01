func findDisappearedNumbers(nums []int) []int {
    missingNums := []int{}
    numSet := make(map[int]bool)
    for i := range nums{
        numSet[nums[i]] = true
    }
    for i := 1; i <= len(nums); i++ {
        if _, ok := numSet[i]; !ok {
            missingNums = append(missingNums, i)
        }
    }
    return missingNums
}
