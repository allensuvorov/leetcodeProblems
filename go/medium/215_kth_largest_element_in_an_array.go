func findKthLargest(nums []int, k int) int {
    // counting sort
    // - 10_000 ... 10_000
    const offset = 10_000
    freqArray := make([]int, 2 * 1e4 + 1)
    for _, v := range nums {
        freqArray[v + offset]++
    }

    ans := 0
    for i := len(freqArray) - 1; i >= 0 && k > 0; i--{
        k -= freqArray[i]
        if k == 0 {
            return i - offset
        }
    }
    return ans    
}
