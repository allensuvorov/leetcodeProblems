func findShortestSubArray(nums []int) int {
    count := map[int]int{}

    degree := 0
    for _, v := range nums {
        count[v]++
        degree = max(degree, count[v])
    }

    ans := len(nums)

    count = map[int]int{}
    l := 0
    for r := range nums {
        count[nums[r]]++
        if count[nums[r]] == degree { // cut the tail
            for l < r && nums[l] != nums[r] {
                count[nums[l]]--
                l++
            }
            ans = min(ans, r - l + 1)
        }
    }
    return ans
}
