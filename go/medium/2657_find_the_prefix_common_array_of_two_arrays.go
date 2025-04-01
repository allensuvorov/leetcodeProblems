func findThePrefixCommonArray(A []int, B []int) []int {
    n := len(A)
    setA := make([]int, n+1) // 0,1,1,1
    setB := make([]int, n+1) // 0,1,1,1
    
    res := make([]int, n) // 0

    for i := range res {
        setA[A[i]]++
        setB[B[i]]++

        if A[i] == B[i] {
            res[i]++    
        } else {
            res[i] += setA[B[i]] + setB[A[i]]
        }

        if i > 0 {
            res[i] += res[i-1]
        }
    }
    return res
}

// idx  0
// A = [2,3,1], 
// B = [3,1,2]

// res  0,1,3



