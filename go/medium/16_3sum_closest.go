func threeSumClosest(nums []int, target int) int {
    slices.Sort(nums)
    smallestDiff := math.MaxInt
    threeSum := 0

    for i, v := range nums {
        l, r := i+1, len(nums)-1
        for l < r {
            sum := v + nums[l] + nums[r]
            if sum == target {
                return sum
            }
        
            if diff := abs(target, sum); diff < smallestDiff {
                smallestDiff = diff
                threeSum = sum
            }

            if sum > target {
                r--
            }

            if sum < target {
                l++
            }
        }
    }
    return threeSum
}

func abs(a, b int) int {
    c := a - b
    if c < 0 {
        return -c
    }
    return c
}
