func getLucky(s string, k int) int {
    sum := 0
    for i := range s {
        num := int(s[i] - 'a' + 1)
        sum += addDigits(num)
    }
    for range k - 1 {
        sum = addDigits(sum)
    }
    return sum
}

func addDigits(num int) int {
    sum := 0
    for x := num; x > 0; x /= 10 {
        digit := x % 10 
        sum += digit
    }
    return sum
}
