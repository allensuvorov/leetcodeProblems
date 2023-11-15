func findGCD(nums []int) int {
    // find min and max
    min, max := 1000, 0
    for _, v := range nums {
        if v > max {
            max = v
        }
        if v < min {
            min = v
        }
    }
    // find max common divisor
    for num := min; num > 1; num-- {
        if max % num == 0 && min % num == 0 {
            return num
        }
    }
    return 1
}
