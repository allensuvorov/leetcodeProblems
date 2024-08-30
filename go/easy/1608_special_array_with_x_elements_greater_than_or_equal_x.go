func specialArray(nums []int) int {
    slices.Sort(nums)
    for i := range nums {
        x := len(nums) - i
        if x <= nums[i] && (i == 0 || nums[i-1] < x){
            return x
            
        }
    }
    return -1
}
