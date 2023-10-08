// (k * (k + 1)) / 2 <= n

/
func arrangeCoins(n int) int {
    return int(math.Sqrt(2.0 * float64(n) + 0.25) - 0.5)
}
*/

func arrangeCoins(n int) int {
    sum := 0
    inc := 1
    i := 0
    for ; sum < n; i++  {
        sum += inc
        inc ++
        if sum > n {
            return i
        }
    }
    return i
}
