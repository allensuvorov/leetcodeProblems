func findThePrefixCommonArray(A []int, B []int) []int {
    n := len(A)
    seen := make([]int, n+1)
    res := make([]int, n)

    for i := range res {

        if A[i] == B[i] {
            res[i]++    
        } else {
            res[i] += seen[B[i]] + seen[A[i]]
        }

        seen[A[i]] |= 1
        seen[B[i]] |= 1

        if i > 0 {
            res[i] += res[i-1]
        }
    }
    return res
}
