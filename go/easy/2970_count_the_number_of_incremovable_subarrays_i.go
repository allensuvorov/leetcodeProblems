func incremovableSubarrayCount(nums []int) int {
    var ans int
    for l := range nums{
        for r := l; r < len(nums); r++ {
            var isIncremovable int = 1 
            temp := append(slices.Clone(nums[:l]), nums[r+1:]...)
            for i := 1; i < len(temp); i++ {
                if temp[i-1] >= temp[i] {
                    isIncremovable = 0
                }
            }
            ans += isIncremovable
        }
    }
    return ans
}
