func decodeString(s string) string {
    stack := make([]byte, 0, len(s))

    for i, v := range s {
        if v != ']' {
            stack = append(stack, s[i])
        } else {
            subStr := []byte{}
            for stack[len(stack) - 1] != '[' {
                subStr = append([]byte{stack[len(stack) - 1]}, subStr...)
                stack = stack[ : len(stack) - 1]
            }
            stack = stack[ : len(stack) - 1]
            numStr := []byte{}
            for len(stack) != 0 && unicode.IsDigit(rune(stack[len(stack) - 1])) {
                numStr = append([]byte{stack[len(stack) - 1]}, numStr...)
                stack = stack[ : len(stack) - 1]
            }
            k, err := strconv.Atoi(string(numStr))
            if err != nil {
		        fmt.Println(err)
	        }
            stack = append(stack, []byte(strings.Repeat(string(subStr), k))...)
        }
    }
    return string(stack)
}
