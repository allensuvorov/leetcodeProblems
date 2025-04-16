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

// 2 hour long and over complicated solution
func calculate(s string) int {
    // space O(1) solution
    // parse the string
    s = strings.ReplaceAll(s, " ", "")
    n := len(s)
    res := 0
    prod := 1
    var prodOp, currOp, nextOp byte = '+', '+', '+'
    for i := 0; i < n; {
        // read number
        num := 0
        for i < n && unicode.IsDigit(rune(s[i])) {
            num *= 10
            num += int(s[i] - '0')
            i++
        }
    
        if i < n {
            nextOp = s[i]
            i++
        }

        // use number
        if currOp == '*' || currOp == '/' { // num is a factor
            prod = doMath(prod, num, currOp) // factor-in the num
            if i == n || nextOp == '+' || nextOp == '-' { 
                res = doMath(res, prod, prodOp) //
            }
        } else if currOp == '+' || currOp == '-' {
            if i == n || nextOp == '+' || nextOp == '-' { // it's a term
                res = doMath(res, num, currOp)
            } else if nextOp == '*' || nextOp == '/' { // new product
                prod = num
                prodOp = currOp
            } 
        }
        
        num = 0 // reset num
        currOp = nextOp // save op to prodOp
    }
    return res
}

func doMath(a, b int, op byte) int {
    switch op {
    case '+': return a + b
    case '-': return a - b
    case '*': return a * b
    case '/': return a / b
    }
    return 0
}
