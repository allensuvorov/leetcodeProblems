func removeDuplicateLetters(s string) string {
    remainingCharCount := [26]int{} // char - 'a'
    for _, v := range s {
        remainingCharCount[v - 'a']++
    }
    
    // increasing stack
    stack := []rune{}
    stackCharCount := [26]int{}

    for _, v := range s {
        if stackCharCount[v - 'a'] == 0 { // do not add duplicates
                    // for something on the stack and
            for len(stack) > 0 && 
                    // top value is greater than v and
                stack[len(stack)-1] > v &&
                    // duplicate we can pop it
                remainingCharCount[stack[len(stack)-1] - 'a'] > 0 { 
                stackCharCount[stack[len(stack)-1] - 'a']--
                stack = stack[:len(stack)-1]
            }
            stackCharCount[v - 'a']++
            stack = append(stack, v)
        }
        remainingCharCount[v - 'a']--
    }
    
    return string(stack)
}
