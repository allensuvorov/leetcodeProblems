func checkPerfectNumber(num int) bool {
    end := int(math.Trunc(math.Sqrt(float64(num))))
    sum := 0
    for i := 1; i <= end; i++ {
        if num%i == 0 {
            if i != num {
                sum += i
            }

            if i != 1 || num/end != end && i == end {
                sum += num/i
            } 
        }
        if sum > num {
            return false
        }
    }
    
    if sum == num {
        return true
    } else {
        return false
    }
}
