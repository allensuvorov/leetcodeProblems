func intersection(nums [][]int) []int {
    a := make([]int, 1001)
    for i := range nums {
        for j := range nums[i] {
            a[nums[i][j]]++
        }
    }
    ans := []int{}
    for i := range a {
        if a[i] == len(nums) {
            ans = append(ans, i)
        }
    }
    return ans
}
