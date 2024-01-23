func countPrimes(n int) int {
    nums := make([]int, n)
    count := 0
    for p := 2; p < len(nums); p++ {
        if nums[p] != -1 {
            count++
            for i := p*p; i < len(nums); i += p {
                nums[i] = -1
            }
        }
    }
    return count
}
