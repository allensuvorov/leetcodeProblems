func maxFrequencyElements(nums []int) int {
    freqCount := map[int]int{}
    maxFreq := 0
    res := 0

    // count freqs
    for _, v := range nums {
        freqCount[v]++
        maxFreq = max(maxFreq, freqCount[v])
    }

    for _, freq := range freqCount {
        if freq == maxFreq {
            res += freq
        }
    }
    return res
}
