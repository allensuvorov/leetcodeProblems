func rotate(nums []int, k int)  {
    l := len(nums)
    // getNew zero index
    z := l - k
    if k > l {
        z = l - (k%l)
    }
    // reslice
    bufNums := append(nums[z:], nums[:z]...)
    copy(nums, bufNums)
}
