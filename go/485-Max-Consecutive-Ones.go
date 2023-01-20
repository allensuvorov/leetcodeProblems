func findMaxConsecutiveOnes(nums []int) int {
    count := 0
    max := 0
    for i, n := range nums {
        if n == 0 && count > 0 {
            if count > max {
                max = count
            }
            count = 0
        }
        if n == 1 {
            count++
            if i == len(nums)-1 && count > max{
                max = count
            }
        }
    }
    return max
}
