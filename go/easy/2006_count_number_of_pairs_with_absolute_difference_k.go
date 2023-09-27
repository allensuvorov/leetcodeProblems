func countKDifference(nums []int, k int) int {
    count := 0
    numCount := map[int]int{}
    for _, v := range nums {
        numCount[v]++
    }
    for i := range numCount {
        count += numCount[i] * numCount[i+k]
    }
    return count
}
