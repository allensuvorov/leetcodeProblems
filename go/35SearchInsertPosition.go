func searchInsert(nums []int, target int) int {
    
    l := len(nums)-1 // biggest index
    mid := math.Floor(l/2)
    
    for {
        if target == nums[mid] {
            return mid
        }

        if target > nums[mid] {
            mid = math.Floor(mid + l-1)/2 // mid to len
        } else {
            mid = math.Floor(mid/2) // 0 to mid
        }
        
        if mid == 0 
    }
    
    return len(nums)
}