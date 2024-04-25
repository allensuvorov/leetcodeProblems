func separateDigits(nums []int) []int {
    ans := make([]int, len(nums)*6)
    r := len(ans) - 1
    for i := len(nums)-1; i >= 0; i-- {
        for tmp := nums[i]; tmp > 0; tmp /= 10 {
            ans[r] = tmp % 10
            r--
        }
    }
    return ans[r+1:]
}
