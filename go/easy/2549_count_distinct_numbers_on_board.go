func distinctIntegers(n int) int {
    onBoard := make(map[int]struct{})
    onBoard[n] = struct{}{}
    
    prev := 0
    for d := 1; d < 10e9 && len(onBoard) != prev; d++ {
        prev = len(onBoard)
        for x := range onBoard {
            for i := 2; i < x; i ++ {
                if x % i == 1 {
                    onBoard[i] = struct{}{}
                }
            }
        }
    }
    return len(onBoard)
}
