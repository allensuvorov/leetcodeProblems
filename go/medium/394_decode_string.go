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

// redo 14 Nov 2025
func decodeString(s string) string {
    stack := []rune{}
    for _, v := range s {
        if v != ']' {
            stack = append(stack, v)
        } else { // decode
            openBracketIndex := len(stack) - 1
            for openBracketIndex >= 0 && stack[openBracketIndex] != '[' {
                openBracketIndex--
            }
            letters := append([]rune{}, stack[openBracketIndex + 1:]...)
            stack = stack[:openBracketIndex]

            // parse number
            k := 0
            position := 1
            for len(stack) > 0 && unicode.IsNumber(stack[len(stack) - 1]) {
                digit := int(stack[len(stack) - 1] - '0')
                digit *= position
                k += digit
                position *= 10
                stack = stack[:len(stack) - 1]
            }

            // k * string
            subString := []rune{}
            for range k {
                subString = append(subString, letters...)
            }

            stack = append(stack, subString...)
        }
        
    }
    return string(stack)
}
