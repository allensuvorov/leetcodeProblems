// GO 1.21
func fib(n int) int {
    a, b := 0, 1
    for i := 1; i <= n; i++ {
        a, b = b, a+b
    }
    return a
}

// GO 1.22
func fib(n int) int {
	a, b := 0, 1
	for _ = range n {
		a, b = b, a+b
	}
	return a
}

