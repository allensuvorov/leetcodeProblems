func calculate(s string) int {
    sum, _ := evaluate([]rune(s))
    return sum
}

func evaluate(s []rune) (int, int) {
    sum, num := 0, 0
    sign := 1
    numReady := false
    for i := 0; i < len(s); i++{
        if !unicode.IsNumber(s[i]) && numReady {
            sum += num * sign
            num, numReady, sign = 0, false, 1
        }
        
        switch s[i] {
        case ' ', '+': // skip
        case '(':
            num, i = evaluate(s[i+1:])
            i = len(s) - i
            sum += num * sign
            num, numReady, sign = 0, false, 1
        case ')':
            return sum, len(s) - i
        case '-':
            sign = -1
        }

        if unicode.IsNumber(s[i]) {
            num = num*10 + int(s[i] - '0')
            numReady = true
        }
    }
    if num > 0 {
        sum += num * sign
    }
    return sum, 0
}
