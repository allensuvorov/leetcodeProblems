// https://www.geeksforgeeks.org/longest-monotonically-increasing-subsequence-size-n-log-n/

func lengthOfLIS(nums []int) int {
    // Binary Search approach
    n := len(nums)
    ans := []int{}
    ans = append(ans, nums[0])
    for i := 1; i < n; i ++ {
        if nums[i] > ans[len(ans)-1] {
            ans = append(ans, nums[i])
        } else {
            // binary search
            l, r := 0, len(ans) - 1
            for l < r {
                mid := l + (r-l)/2
                if ans[mid] < nums[i] {
                    l = mid + 1
                } else {
                    r = mid
                }
            }
            ans[l] = nums[i]
        }
    }
    return len(ans)
}
