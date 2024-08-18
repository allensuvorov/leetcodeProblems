func minChanges(nums []int, k int) int {
    if k == 0 {
        return 0
    }
    ans := math.MaxInt
    diffFrequencyCount := make(map[int]int)
    singleChangeCapacityFrequencyCount := make([]int, k + 1)

    for l, r := 0, len(nums) - 1; l < r; l, r = l+1, r-1 {
        absDiff := absDiff(nums[l], nums[r])
        diffFrequencyCount[absDiff]++
        singleChangeCapacity := max(nums[l], nums[r], k - nums[l], k - nums[r])
        singleChangeCapacityFrequencyCount[singleChangeCapacity]++
    }
    
    prefixSum := make([]int, k + 1)
    for sum, i := 0, k; i >= 0; i-- {
        v := singleChangeCapacityFrequencyCount[i]
        sum += v
        prefixSum[i] = sum
    }

    for diff, freqCount := range diffFrequencyCount {
        singleChangeCount := (prefixSum[diff] - freqCount)
        doubleChangeCount := len(nums)/2 - singleChangeCount - freqCount
        totalChangeCount := singleChangeCount + doubleChangeCount * 2
        ans = min(ans, totalChangeCount)
    }
    return ans

}

func absDiff(a, b int) int {
    c := a - b
    if c < 0 {
        c = -c
    }
    return c
}
