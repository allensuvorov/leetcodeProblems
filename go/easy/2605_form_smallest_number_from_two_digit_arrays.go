func minNumber(nums1 []int, nums2 []int) int {
    min1 := Min(nums1)
    min2 := Min(nums2)

    n, ok := minCommon(nums1, nums2)
    if ok {
        return n
    }

    switch {
    case min1 > min2:
        return min2*10 + min1
    case min1 < min2:
        return min1*10 + min2
    }
    return min1
}

func minCommon(nums1 []int, nums2 []int) (int, bool) {
    b1, b2 := 0, 0
    for i := range nums1 {
        b1 |= 1 << nums1[i]
    }
    for i := range nums2 {
        b2 |= 1 << nums2[i]
    }
    b := b1 & b2
    
    if b == 0 {
        return 0, false
    }
    return bits.TrailingZeros(uint(b)), true
}

func Min(nums []int) int {
    min := 10
    for i := range nums {
        if min > nums[i] {
            min = nums[i]
        }
    }
    return min
}
