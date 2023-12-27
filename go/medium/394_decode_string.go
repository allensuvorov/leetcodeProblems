func decodeString(s string) string {
    // Transform from in-fix to post-fix expression
    postFixExpression := buildPostFixExpression(s)
    fmt.Printf("%s", postFixExpression)

    // Calculate the sum
    res := getSum(postFixExpression)
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
        // fmt.Printf("loop start: v = %v, postFixExpression = %s \n", string(v), postFixExpression)
        if unicode.IsDigit(v) {
            digits = append(digits, s[i])
            if len(letters) != 0 {
                // fmt.Printf("about to flush letter to pfe: %s \n", postFixExpression)
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
            data := make([]byte, len(letters))
            copy(data, letters)
            postFixExpression = append(postFixExpression, data)
            letters = letters[:0]
            for stack[len(stack)-1] != '[' {
                popFromStackToPostFixExpression()
            }
            stack = stack[:len(stack)-1]

            if i + 1 < len(s) && s[i + 1] != ']' {
                pushPlus()
            }
        }
        // fmt.Printf("loop end: v = %v, postFixExpression = %s \n", string(v), postFixExpression)
    }
    for len(stack) != 0 {
        popFromStackToPostFixExpression()
    }
    return postFixExpression
}

func getSum(pfe [][]byte) []byte {
    for i, v := range pfe {
        // what is it? numbers or letters
        switch {
        case unicode.IsDigit(v[0]):
        case unicode.IsLetter(v[0]):
        case v[0] == '+':
        case v[0] == '*':
        }
    }
}

// in-fix: 3 * a + 2 * bc
// post-fix: 3 a * 2 bc * +


// in-fix: 3 * [a + 2 * [c] ] = 3 * [a + cc]
// post-fix: 3a2c*+*
