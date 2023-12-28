func decodeString(s string) string {
    postFixExpression := buildPostFixExpression(s)
    fmt.Printf("%s", postFixExpression)
    res := getCalculation(postFixExpression)
    return string(res[0])
}

func buildPostFixExpression(s string) [][]byte{
    stack := make([]byte, 0, len(s))
    postFixExpression := [][]byte{}
    digits := []byte{}
    letters := []byte{}
    popFromStackToPostFixExpression := func() {
        postFixExpression = append(postFixExpression, []byte{stack[len(stack)-1]})
        stack = stack[:len(stack)-1]
    }
    pushPlus := func() {  // "+" pushes all "+" and "*" from stack to post-fix-exp
        for len(stack) > 0 && (stack[len(stack)-1] == '+' || stack[len(stack)-1] == '*') {
            popFromStackToPostFixExpression()
        }
        stack = append(stack, '+')    
    }
    for i, v := range s {
        if unicode.IsDigit(v) {
            digits = append(digits, s[i])
            if len(letters) != 0 {
                data := make([]byte, len(letters))
                copy(data, letters)
                postFixExpression = append(postFixExpression, data)
                letters = letters[:0]
                pushPlus()
            }
        }
        if unicode.IsLetter(v) {
            letters = append(letters, s[i])
        }
        if v == '['{
            data := make([]byte, len(digits))
            copy(data, digits)
            postFixExpression = append(postFixExpression, data)
            digits = digits[:0]
            for len(stack) > 0 && stack[len(stack)-1] == '*' {
                popFromStackToPostFixExpression()
            }
            stack = append(stack, '*')
            stack = append(stack, '[')
        }
        if v == ']'{
            if len(letters) != 0 {
                data := make([]byte, len(letters))
                copy(data, letters)
                postFixExpression = append(postFixExpression, data)
                letters = letters[:0]
            }
            for stack[len(stack)-1] != '[' {
                popFromStackToPostFixExpression()
            }
            stack = stack[:len(stack)-1]
            if i + 1 < len(s) && s[i + 1] != ']' {
                pushPlus()
            }
        }
    }
    if len(letters) != 0 {
        data := make([]byte, len(letters))
        copy(data, letters)
        postFixExpression = append(postFixExpression, data)
        letters = letters[:0]
    }
    for len(stack) != 0 {
        popFromStackToPostFixExpression()
    }
    return postFixExpression
}

func getCalculation(pfe [][]byte) [][]byte {
    stack := make([][]byte, 0, len(pfe))
    for _, v := range pfe {
        switch {
        case unicode.IsDigit(rune(v[0])):
            stack = append(stack, v)
        case unicode.IsLetter(rune(v[0])):
            stack = append(stack, v)
        case v[0] == '+':
            b := stack[len(stack) - 1]
            a := stack[len(stack) - 2]
            stack = stack[ : len(stack) - 2]
            c := append(a, b...)
            stack = append(stack, c)
        case v[0] == '*':
            letters := stack[len(stack) - 1]
            k := buildNum(stack[len(stack) - 2])
            repeated := repeat(k, letters)
            stack = stack[ : len(stack) - 2]
            stack = append(stack, repeated)
        }
    }
    return stack
}

func buildNum(digits []byte) int {
    num := 0
    for _, v := range digits {
        num = num * 10 + int(v - '0')
    }
    return num
}

func repeat(k int, letters []byte) []byte {
    res := make([]byte, 0, k * len(letters))
    for i := 1; i <= k; i++ {
        res = append(res, letters...)
    }
    return res
}
