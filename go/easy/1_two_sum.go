func twoSum(nums []int, target int) []int {
    prevNums := make(map[int]int)
    for i, curNum := range nums {
        if j, ok := prevNums[target - curNum]; ok {
            return []int{i, j}
        }
        prevNums[curNum] = i
    }
    return nil
}
