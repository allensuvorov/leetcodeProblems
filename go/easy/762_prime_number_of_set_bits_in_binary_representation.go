func countPrimeSetBits(left int, right int) int {
    primes := map[int]bool{2:true, 3:true, 5:true, 7:true, 11:true, 13:true, 17:true, 19:true}
    // indecies are numbers
    bitCount := make([]int, right + 1)
    ans := 0
    for i := 1; i <= right; i++ {
        // 1 probrem: count bits
        if i & 1 == 0 {
            bitCount[i] = bitCount[i/2]
        } else {
            bitCount[i] = bitCount[i-1] + 1
        }
        // 2 problem: check for prime
        if i >= left && primes[bitCount[i]]{
            ans++
        }
    }
    return ans
}
