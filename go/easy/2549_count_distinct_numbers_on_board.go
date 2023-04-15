func distinctIntegers(n int) int {
    onBoard := make(map[int]struct{})
    onBoard[n] = struct{}{}
    
    prev := 0
    for d := 1; d < 10e9 && len(onBoard) != prev; d++ {
        prev = len(onBoard)
        for x := range onBoard {
            for i := 2; i < n; i ++ { // 2 - 4
                if x % i == 1 { // 5 
                    onBoard[i] = struct{}{} // 2, 4
                }
            }
        }
    }
    return len(onBoard)
}
