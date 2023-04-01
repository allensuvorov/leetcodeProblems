func sortArrayByParityII(nums []int) []int {
    even := 0
    odd := 1
    for even <= len(nums)-2 && odd <= len(nums)-1 {
        if nums[even]%2 == 0 {
            even += 2 
            continue
        }
        if nums[odd]%2 != 0 {
            odd += 2
            continue
        }
        nums[even], nums[odd] = nums[odd], nums[even]
    }
    return nums
}
