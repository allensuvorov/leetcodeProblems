func twoSum(nums []int, target int) []int {
    numIndex := make(map[int]int, len(nums))
    for i, v := range nums {
        if j, ok := numIndex[target - v]; ok {
            return []int{i, j}
        }
        numIndex[v] = i
    }
    return nil
}
