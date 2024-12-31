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

// redo
func decodeString(s string) string {
	stack := []string{}
	for _, v := range s {
		switch {
		case unicode.IsLetter(v):
			if len(stack) == 0 || stack[len(stack)-1] == "[" {
				stack = append(stack, string(v))
			} else {
				stack[len(stack)-1] += string(v)
			}
		case unicode.IsDigit(v):
			if len(stack) == 0 || !unicode.IsNumber(rune(stack[len(stack)-1][0])) {
				stack = append(stack, string(v))
			} else {
				stack[len(stack)-1] += string(v)
			}
		case v == '[':
			stack = append(stack, string(v))
		case v == ']':
			word := stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			number, err := strconv.Atoi(stack[len(stack)-1])
			if err != nil {
				fmt.Println(err)
			}
			stack = stack[:len(stack)-1]
            		multiWord := strings.Repeat(word, number)
			if len(stack) > 0 && unicode.IsLetter(rune(stack[len(stack)-1][0])) {
				stack[len(stack)-1] += multiWord
			} else {
				stack = append(stack, multiWord)
			}
		}
	}
	return stack[0]
}
