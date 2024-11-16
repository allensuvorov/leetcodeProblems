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
            sum += sign * num
            num, numReady = 0, false
        }

        switch s[i] {
        case ' ': // skip
        case '+':
            sign = 1
        case '-':
            sign = -1
        case '(':
            num, i = evaluate(s[i+1:])
            i = len(s) - i
            sum += sign * num
            num, numReady = 0, false
        case ')':
            return sum, len(s) - i
        }

        if unicode.IsNumber(s[i]) {
            num = num*10 + int(s[i] - '0')
            numReady = true
        }
    }
    if num > 0 {
        sum += sign * num
    }
    return sum, 0
}
