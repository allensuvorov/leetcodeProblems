func minMaxDifference(num int) int {
    digits := []int{}
    for x := num; x > 0; x /= 10 {
        digits = append([]int{x%10}, digits...)
    }

    target := -1
    for i := range digits {
        if digits[i] != 9 {
            target = digits[i]
            break
        }
    }
    max := num
    if target != -1 {
        max = 0
        for i := range digits {
            d := digits[i]
            if d == target {
                d = 9
            }
            max = max*10 + d
        }
    }
    min := 0
    target = digits[0]
    for i := 1; i < len(digits); i++ {
        d := digits[i]
        if d == target {
            d = 0
        }
        min = min*10 + d
    }
    return max - min
}
