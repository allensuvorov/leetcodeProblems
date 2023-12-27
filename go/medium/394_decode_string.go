func decodeString(s string) string {
    // Transform from in-fix to post-fix expression
    postFixExpression := buildPostFixExpression(s)
    fmt.Printf("%s", postFixExpression)

    // Calculate the sum
        // what is it? numbers or letters
        return ""
}


func buildPostFixExpression(s string) [][]byte{
    stack := make([]byte, 0, len(s))
    postFixExpression := [][]byte{}
    digits := []byte{}
    letters := []byte{}
    for i, v := range s {
        fmt.Printf("loop start: v = %v, postFixExpression = %s \n", string(v), postFixExpression)
        if unicode.IsDigit(v) {
            digits = append(digits, s[i])
            if len(letters) != 0 {
                fmt.Printf("about to flush letter to pfe: %s \n", postFixExpression)
                data := make([]byte, len(letters))
                copy(data, letters)
                postFixExpression = append(postFixExpression, data)
                letters = letters[:0]
                for len(stack) > 0 && (stack[len(stack)-1] == '+' || stack[len(stack)-1] == '*') {
                    postFixExpression = append(postFixExpression, []byte{stack[len(stack)-1]})
                    stack = stack[:len(stack)-1]
                }
                stack = append(stack, '+')
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
                postFixExpression = append(postFixExpression, []byte{'*'})
                stack = stack[:len(stack)-1]
            }
            stack = append(stack, '*')
            stack = append(stack, '[')
        }
        if v == ']'{
            postFixExpression = append(postFixExpression, letters)
            letters = letters[:0]
            for stack[len(stack)-1] != '[' {
                postFixExpression = append(postFixExpression, []byte{stack[len(stack)-1]})
                stack = stack[:len(stack)-1]
            }
            stack = stack[:len(stack)-1]
        }
        fmt.Printf("loop end: v = %v, postFixExpression = %s \n", string(v), postFixExpression)
    }
    for len(stack) != 0 {
        postFixExpression = append(postFixExpression, []byte{stack[len(stack)-1]})
        stack = stack[:len(stack)-1]
    }
    return postFixExpression
}


// 3 * [a + 2 * [c] ] = 3 * [a + cc]

// post-fix: 3a2c*+*
// stack: 3a2c
// res:a cc

