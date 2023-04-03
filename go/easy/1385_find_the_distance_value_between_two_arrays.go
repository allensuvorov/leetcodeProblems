func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
    n := 0
    for i := range arr1 {
        ok := true
        for j := range arr2 {
            absDiff := arr1[i] - arr2[j]
            if absDiff <= 0 {
                absDiff *= -1
            }
            
            if absDiff <= d {
                ok = false
                break
            }
        }
        if ok {
            n++
        }
    }
    return n
}
