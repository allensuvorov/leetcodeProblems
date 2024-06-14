func calculate(s string) int {
    curNum, lastNum, result := 0, 0, 0
    operation := '+'
    for i, curChar := range s {
        if unicode.IsDigit(curChar) {
            curNum = (curNum * 10) + int(curChar - '0')
        }
        if !unicode.IsDigit(curChar) && !unicode.IsSpace(curChar) || i == len(s) - 1 {
            switch operation {
            case '*':
                lastNum *= curNum
            case '/':
                lastNum /= curNum
            case '+':
                result += lastNum
                lastNum = curNum
            case '-':
                result += lastNum
                lastNum = -curNum
            }
            operation = curChar
            curNum = 0
        }
    }
    result += lastNum
    return result
}
