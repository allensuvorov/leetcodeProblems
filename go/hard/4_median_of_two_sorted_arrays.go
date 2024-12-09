func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    a, b := nums1, nums2
    if len(a) > len(b) {
        a, b = b, a // Ensure the smaller array is 'a'
    }

    total := len(a) + len(b)
    half := total / 2

    l, r := 0, len(a)

    for l <= r {
        i := l + (r-l)/2 // Partition index for 'a'
        j := half - i    // Partition index for 'b'

        aLeft := math.MinInt64
        if i > 0 {
            aLeft = a[i-1]
        }

        aRight := math.MaxInt64
        if i < len(a) {
            aRight = a[i]
        }

        bLeft := math.MinInt64
        if j > 0 {
            bLeft = b[j-1]
        }

        bRight := math.MaxInt64
        if j < len(b) {
            bRight = b[j]
        }

        // Check for correct partition
        if aLeft <= bRight && bLeft <= aRight {
            // Odd total length
            if total%2 == 1 {
                return float64(min(aRight, bRight))
            }
            // Even total length
            return float64(max(aLeft, bLeft)+min(aRight, bRight)) / 2.0
        } else if aLeft > bRight {
            r = i - 1 // Move left in 'a'
        } else {
            l = i + 1 // Move right in 'a'
        }
    }

    return 0.0 // This will never be reached if input is valid
}
