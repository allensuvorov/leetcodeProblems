func twoSum(nums []int, target int) []int {
    m := map[int]int{}
    for i, a := range nums {
        b := target - a
        if j, ok := m[b]; ok {
            return []int{j, i}
        }
        m[a] = i
    }
    return nil
}
