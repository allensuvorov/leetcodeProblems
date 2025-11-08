// count sort with smaller n
func frequencySort(nums []int) []int {
    const offset = 100
    freqList := make([]int, 201)
    
    for _, v := range nums {
        freqList[v + offset]++
    }

    freqTable := make([][]int, 101)
    for num, freq := range freqList {
        if freq > 0 {            
            freqTable[freq] = append(freqTable[freq], num - offset)
        }
    }

    ans := []int{}
    for freq, nums := range freqTable {
        for i := len(nums) - 1; i >= 0; i-- {
            for range freq {
                ans = append(ans, nums[i])
            }
        }
    }
    return ans
}

// general case - with hash map
func frequencySort(nums []int) []int {
    freqCount := map[int]int{}
    for _, v := range nums {
        freqCount[v]++
    }

    freqs := []freqRecord{}
    for val, freq := range freqCount {
        freqs = append(freqs, freqRecord{
            val: val,
            freq: freq,
            })
    }

    slices.SortFunc(freqs, 
        func(a, b freqRecord) int {
            return cmp.Or( // return first non zero
                cmp.Compare(a.freq, b.freq),
                cmp.Compare(b.val, a.val),
            )
        },
    )
    
    res := make([]int, 0, len(nums))
    for _, v := range freqs {
        for range v.freq {
            res = append(res, v.val)
        }
    }
    return res
}

type freqRecord struct {
    val int
    freq int
}
