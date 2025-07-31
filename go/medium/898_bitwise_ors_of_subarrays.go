func subarrayBitwiseORs(arr []int) int {
    orSet := make(map[int]bool)

    for i := range arr {
        or := 0
        for j := i; j < len(arr); j++ {
            or |= arr[j]
            orSet[or] = true
        }
    }

    return len(orSet)
}

// 1
// 11
// 112
//  1
//  12
//   2
