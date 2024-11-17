/*
* Big O (n) Time
* Big O (1) Space
* Two pointers
* Math
*/

func incremovableSubarrayCount(nums []int) int64 {
    ans := 0
    n := len(nums)

    leftIncIdx := 0 // locate increasing region on the left side
    for i := 0; i < n - 1 && nums[i] < nums[i+1]; i++ {
        leftIncIdx = i + 1
    }
    // if fully increasing - do math
    if leftIncIdx == n - 1 {
        return int64(n * (n + 1) / 2)
    }

    rightIncIdx := n - 1  // locate increasing region on the right side
    for i := n - 1; i > 0 && nums[i] > nums[i-1]; i-- {
        rightIncIdx = i - 1
    }
    // count empty array and one-side increasing regions separately
    empty, leftSide, rightSide := 1, leftIncIdx + 1, n - rightIncIdx
    ans += empty + leftSide + rightSide
    // for every val on the left find increasing reagion on the right
    for l, r := 0, rightIncIdx; l <= leftIncIdx && r < n; l++ {        
        for r < n && nums[l] >= nums[r] {
            r++
        }
        if r < n && nums[l] < nums[r] {
            ans += (n-r)
        }
    }
    return int64(ans)
}
