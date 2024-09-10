func maximumSwap(num int) int {
    stack := []int{}
    maxDigVal := 0
    maxDigPos := 0
    position := 0
    
    for x := num; x > 0; x /= 10 {
        digit := x % 10
        if digit > maxDigVal {
            maxDigVal = digit
            maxDigPos = position
        }
        stack = append(stack, digit)
        position++
    }
    minDigPos := 0
    for i := len(stack) - 1; i > maxDigPos; i-- {
        if stack[i] < maxDigVal {
            minDigPos = i
        }
    }
    if minDigPos > 0 {
        stack[minDigPos], stack[maxDigPos] = stack[maxDigPos], stack[minDigPos]
    }
    ans := 0
    for i := len(stack) - 1; i >= 0; i-- {
        ans *= 10
        ans += stack[i]
    }
    return ans

}

// 2736

// 66177
// 76176

// 1234
// stack: 4321



