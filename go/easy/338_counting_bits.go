func countBits(n int) []int {
    ans := make([]int, n + 1)
    for i := 1; i <= n; i++ {        
        if i & 1 == 0 { // even
            ans[i] = ans[i/2]
        } else { // odd
            ans[i] = ans[i-1] + 1
        }
    }
    return ans
}


// Big O (n lon n)

func countBits(n int) []int {
    ans := make([]int, n + 1)
    for num := range n + 1 {
        for x := num; x > 0; x >>= 1 {
            ans[num] += x & 1
        }
    }
    return ans
}
