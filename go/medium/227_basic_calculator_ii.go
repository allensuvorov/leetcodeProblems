func calculate(s string) int {
    n := len(s)
    res, prod, num := 0, 0, 0
    op := '+'
    for i, c := range s {
        // build number
        if unicode.IsDigit(c) {
            num = (num * 10) + int(c - '0')
        } 
        // do operation
        if i == n - 1 || c == '+' || c == '-' || c == '*' || c == '/' {
            switch op {
            case '+':
                res += prod
                prod = num
            case '-':
                res += prod
                prod = -num
            case '*':
                prod *= num
            case '/':
                prod /= num
            }
            op = c
            num = 0
        }
    }
    res += prod
    return res
}
