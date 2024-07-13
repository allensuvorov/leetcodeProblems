func sumOfTheDigitsOfHarshadNumber(x int) int {
    sum := 0
    for y := x; y > 0; y /= 10 {
        digit := y%10
        sum += digit
    }

    if x % sum == 0 {
        return sum
    }
    return -1
}
