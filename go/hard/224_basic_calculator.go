func calculate(s string) int {
    sum, _ := getSum([]rune(s))
    return sum
}

func getSum(s []rune) (int, int) {
    sum := 0
    num := 0
    minus := false
    numExists := false
    for i := 0; i < len(s); i++{
        if !unicode.IsNumber(s[i]) && numExists {
            if minus {
                sum -= num
            } else {
                sum += num
            }
            num = 0
            numExists = false
            minus = false
        }

        if unicode.IsNumber(s[i]) {
            num = num*10 + int(s[i] - '0')
            numExists = true
        }

        switch s[i] {
        case ' ', '+':
            // skip
        case '(':
            num, i = getSum(s[i+1:])
            i = len(s) - i
            if minus {
                sum -= num
            } else {
                sum += num
            }
            num = 0
            numExists = false
            minus = false
        case ')':
            return sum, len(s) - i
        case '-':
            minus = true
        }
    }
    if num > 0 {
        if minus {
            sum -= num
        } else {
            sum += num
        }
    }
    return sum, 0
}
