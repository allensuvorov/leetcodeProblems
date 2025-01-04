func twoSum(nums []int, target int) []int {
    prevs := make(map[int]int)
    for i, cur := range nums {
        if j, ok := prevs[target - cur]; ok {
            return []int{i, j}
        }
        prevs[cur] = i
    }
    return nil
}
