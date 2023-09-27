func countKDifference(nums []int, k int) int {
    if len(nums) == 1 {
        return 0
    }
    count := 0
    numCount := [101]int{}
    for _, v := range nums {
        numCount[v]++
    }
    for i := 1; i <= 100 - k; i++ {
        count += numCount[i] * numCount[i+k]
    }
    return count
}
