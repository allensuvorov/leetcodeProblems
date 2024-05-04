func generateTheString(n int) string {
    ans := make([]byte, n)
    for i := range n {
        ans[i] = 'a'
    }
    if n % 2 == 0 {
        ans[0] = 'b'
    }
    return string(ans)
}
