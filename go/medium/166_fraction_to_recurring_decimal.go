func fractionToDecimal(numerator int, denominator int) string {
    if numerator % denominator == 0 {
        return strconv.Itoa(numerator / denominator)
    }

    sign := ""

    if numerator < 0 && denominator > 0 || numerator > 0 && denominator < 0 {
        sign = "-"
    }
    
    numerator = abs(numerator)
    denominator = abs(denominator)

    whole := numerator / denominator
    rem := numerator % denominator
    seen := map[int]int{}
    fraction := []byte{}

    for rem != 0 {
        seen[rem] = len(fraction)
        rem = rem * 10
        digit := rem / denominator
        rem = rem % denominator
        fraction = append(fraction, byte(digit) + '0')
        
        if index, ok := seen[rem]; ok {
            repeating := append([]byte{'('}, fraction[index:]...)
            repeating = append(repeating, ')')
            fraction = append(fraction[:index], repeating...)
            break
        }
    }

    return sign + strconv.Itoa(whole) + "." + string(fraction)
}

func abs(a int) int {
    if a < 0 {
        a = -a
    }
    return a
}
