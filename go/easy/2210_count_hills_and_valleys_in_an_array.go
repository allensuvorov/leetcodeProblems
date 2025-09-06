func countHillValley(nums []int) int {
    lastDir := 0
    currDir := 0
    count := 0
    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[i-1] {
            currDir = 1
        } 

        if nums[i] < nums[i-1]{
            currDir = -1
        }

        if lastDir != 0 && lastDir == -currDir {
            count++
        }

        lastDir = currDir
    }
    return count
}
