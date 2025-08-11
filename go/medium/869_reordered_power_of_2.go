func reorderedPowerOf2(n int) bool {
    // try each permutation - Big O((Log(10)N)!)
    digits := make([]int, 0)
    for x := n; x > 0; x = x / 10 {
        digits = append(digits, x % 10)
    }
    return dfs(digits, 0)
}


func dfs(digits []int, num int) bool {
    if len(digits) == 0 {
        return isPowerOf2(num)
    }
    
    for i, digit := range digits {
        if digit != 0 || num != 0  {
            nextDigits := make([]int, 0, len(digits) - 1)
            for j, v := range digits {
                if j != i {
                    nextDigits = append(nextDigits, v)
                }
            }
            if dfs(nextDigits, num * 10 + digit) {
                return true
            }
        }
    }

    return false
}

func isPowerOf2(n int)bool {
    return n > 0 && (n & (n-1) == 0)
}
