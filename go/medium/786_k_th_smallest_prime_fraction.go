func kthSmallestPrimeFraction(arr []int, k int) []int {
    fractions := [][]int{}

    for i := range arr {
        for j := i + 1; j < len(arr); j++ {
            fractions = append(fractions, []int{arr[i], arr[j]})
        }
    }

    slices.SortFunc(fractions, func(a, b []int) int {
        diff := float64(a[0])/float64(a[1]) - float64(b[0])/float64(b[1])
        
        if diff > 0.0 {
            return 1
        }
        if diff < 0.0 {
            return -1
        }
        
        return 0
    })

    return fractions[k-1]
}
