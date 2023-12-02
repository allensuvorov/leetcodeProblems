func differenceOfSum(nums []int) int {
    elmSum, dgtSum := 0, 0

    for _, v := range nums {
        elmSum += v
        dgtSum += getDgtSum(v)
    }

    dif := elmSum - dgtSum
    if dif < 0 {
        return - dif
    }
    return dif
}

func getDgtSum(n int) int {
    dgtSum := 0
    for n > 0 {
        d := n % 10
        dgtSum += d
        n /= 10
    }
    return dgtSum
}
