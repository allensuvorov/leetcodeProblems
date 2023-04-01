func kLengthApart(nums []int, k int) bool {
    prev :=-1
    for i := range nums {
        if nums[i] == 1 {
            if prev >= 0 && i - prev < k+1 {
                return false
            }
        prev = i
        }
    }
    return true
}
