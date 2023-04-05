func minimizeArrayValue(nums []int) int {
	prefixSum := 0
	max := 0

	for i, num := range nums {
		prefixSum += num
        if (prefixSum + i) / (i + 1) > max {
            max = (prefixSum + i) / (i + 1)
        }
	}
	return max
}

// following solution ran out of time limit with bigger cases

// func minimizeArrayValue(nums []int) int {    
//     max := 0
//     count := 0
//     for {
//         max = 0
//         for i := range nums {
//             if nums[i] > nums[max] {
//                 max = i
//             } 
//         }

//         if nums[max] == 0 || max == 0 {
//             return nums[max]
//         }

//         diff := nums[max] - nums[max-1]
//         if diff < 2 {
//             nums[max] --
//             nums[max-1] ++
//         } else {
//             nums[max] -= diff/2
//             nums[max-1] += diff/2
//         }
        
//         count ++
//         if count > 10e5 {
//             fmt.Println(count)
//         }
//         if count > 2e6 {
//             return count
//         }

//     }
//     return nums[max]
// }
