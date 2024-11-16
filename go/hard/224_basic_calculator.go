func calculate(s string) int {
    sum, _ := getSum([]rune(s))
    return sum
}

func getSum(s []rune) (int, int) {
    sum, num := 0, 0
    minus := 1
    numExists := false
    for i := 0; i < len(s); i++{
        if !unicode.IsNumber(s[i]) && numExists {
            sum += num * minus
            num, numExists, minus = 0, false, 1
        }

        if unicode.IsNumber(s[i]) {
            num = num*10 + int(s[i] - '0')
            numExists = true
        }

        switch s[i] {
        case ' ', '+': // skip
        case '(':
            num, i = getSum(s[i+1:])
            i = len(s) - i
            sum += num * minus
            num, numExists, minus = 0, false, 1
        case ')':
            return sum, len(s) - i
        case '-':
            minus = -1
        }
    }
    if num > 0 {
        sum += num * minus
    }
    return sum, 0
}
