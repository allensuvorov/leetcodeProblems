func countDigits(num int) int {
    count := 0
    value := num
    for value>0 {
        digit := value%10
        if digit != 0 && num % digit == 0 {
            count++
        } 
        value /= 10
    }
    return count
}
