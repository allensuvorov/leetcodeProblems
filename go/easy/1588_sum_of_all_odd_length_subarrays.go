func sumOddLengthSubarrays(arr []int) int {
    // n^2 with sliding window
    totalSum := 0
    for window := 1; window <= len(arr); window += 2 {
        windowSum := 0
        for i := range arr {
            windowSum += arr[i]
            if i + 1 > window { // cut tailing value
                windowSum -= arr[i - window]
            }
            if i + 1 >= window {
                totalSum += windowSum
            } 
        }
    }
    return totalSum
}

//     1234578
// 1:  1111111
//     abcdefg
// 3:  1233321
//     abc
//      bcd
//       cde
//        def
//         efg
// 5:  1233321
//     abcde
//      abcde
//       abcde
// 7:  1222221
//     abcdef
//      bcdefg
//     1234578
