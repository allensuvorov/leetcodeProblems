func findMaxConsecutiveOnes(nums []int) int {
    count := 0
    max := 0
    for _, n := range nums {
        if n == 1 {
            count++
            if count > max{
                max = count
            }
        } else {
            count = 0
        }
    }
    return max
}
