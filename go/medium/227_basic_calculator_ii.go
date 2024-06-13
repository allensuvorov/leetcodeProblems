func calculate(s string) int {
    if len(s) == 0 {
        return 0
    }

    curNum, lastNum, result := 0, 0, 0
    operation := '+'
    for i, v := range s {
        curChar := v
        if unicode.IsDigit(curChar) {
            curNum = (curNum * 10) + int(curChar - '0')
        }
        if !unicode.IsDigit(curChar) && !unicode.IsSpace(curChar) || i == len(s) - 1 {
            switch operation {
            case '+':
                result += lastNum
                lastNum = curNum
            case '-':
                result += lastNum
                lastNum = -curNum
            case '*':
                lastNum *= curNum
            case '/':
                lastNum /= curNum
            }
            operation = curChar
            curNum = 0
        }
    }
    result += lastNum
    return result
}
