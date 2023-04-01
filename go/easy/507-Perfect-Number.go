func checkPerfectNumber(num int) bool {
    end := int(math.Trunc(math.Sqrt(float64(num))))
    if num == 1 {
        return false
    }
    sum := 0
    for i := 1; i*i <= num; i++ {
        if num%i == 0 {
            sum += i
            if i != 1 || num/end != end && i == end {
                sum += num/i
            } 
        }
        if sum > num {
            return false
        }
    } 
    return sum == num
}
