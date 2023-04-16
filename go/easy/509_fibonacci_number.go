package easy

func fib(n int) int {
    // loop
    a, b := 0, 1
    for i := 1; i <= n; i++ {
        a, b = b, a+b
    }
    return a
}

// other solutions: 
// * closure
// * recursion
