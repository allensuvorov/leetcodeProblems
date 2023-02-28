func relativeSortArray(arr1 []int, arr2 []int) []int {
    hm := map[int]int{}
    for i := range arr1 {
        hm[arr1[i]]++
    }
    var i, j int
    for i < len(arr1) && j < len(arr2) {
        arr1[i] = arr2[j]
        i++
        hm[arr2[j]]--
        if hm[arr2[j]] == 0 {
            delete(hm, arr2[j])
            j++
        }
    }
    minKey := func(m map[int]int) int {
        min := 1000
            for k:= range m {
                if k < min {
                    min = k
                }
            }
        return min
    }
    min := minKey(hm)
    for len(hm)>0 {
        arr1[i] = min
        i++
        hm[min]--
        if hm[min] == 0 {
            delete(hm, min)
            min = minKey(hm)
        }
    }
    return arr1
}
