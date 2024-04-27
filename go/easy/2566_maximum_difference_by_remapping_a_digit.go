func minMaxDifference(num int) int {
    stack := []int{}
    for x := num; x > 0; x /= 10 {
        d := x%10
        stack = append(stack, d)
    }
    target := -1
    for i := len(stack) - 1; i >= 0; i-- {
        if stack[i] != 9 {
            target = stack[i]
            break
        }
    }
    max := num
    if target != -1 {
        max = 0
        for i := len(stack) - 1; i >= 0; i-- {
            d := stack[i]
            if d == target {
                d = 9
            }
            max = max*10 + d
        }
    }
    min := 0
    target = stack[len(stack) - 1]
    for i := len(stack) - 2; i >= 0; i-- {
        d := stack[i]
        if d == target {
            d = 0
        }
        min = min*10 + d
    }
    fmt.Println(stack, max, min)
    return max - min
}
