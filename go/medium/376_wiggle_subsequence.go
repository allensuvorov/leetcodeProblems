//  time O(n)
// space O(1)

func wiggleMaxLength(nums []int) int {
    res := 1
    prev := 0 // 1, -1
    for i := 1; i < len(nums); i++ {
        diff := nums[i] - nums[i-1]
        
        if prev == 0 && diff != 0 {
            if diff > 0 {
                prev = 1
            }  
            if diff < 0 {
                prev = -1
            }
            res++
        }

        if diff * prev < 0 {
            res++
            prev = -prev
        }
    }
    return res
}
