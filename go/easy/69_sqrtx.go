func mySqrt(x int) int {

    l, r := 0, x+1
    
    for l < r {
        
        m := (l + r)/2
        guess := m*m
        fmt.Println(x, guess, l , m, r)

        switch {
            case guess == x: return m
            case guess > x: r = m
            case guess < x: l = m + 1
        }
    }
    return l - 1 
}
