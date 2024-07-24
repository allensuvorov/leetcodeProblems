func sortJumbled(mapping []int, nums []int) []int {
    mappedValues := make([][]int, len(nums))
    
    for i := range nums {
        mv := 0
        decimal := 1
        for x := nums[i]; x > 0; x /= 10 {
            d := x % 10
            md := mapping[d]
            mv += md * decimal
            decimal *= 10
        }
        if nums[i] == 0 {
            mv = mapping[nums[i]]
        }
        mappedValues[i] = []int{mv, nums[i]}
    }

    slices.SortStableFunc(mappedValues, func(a, b []int) int {
        return cmp.Compare(a[0], b[0])
    })


    ans := make([]int, len(nums))
    for i := range mappedValues {
        ans[i] = mappedValues[i][1]
    }
    return ans
}
