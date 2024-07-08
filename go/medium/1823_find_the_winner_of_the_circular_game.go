func findTheWinner(n int, k int) int {
    return winnerHelper(n, k) + 1
}

func winnerHelper(n int, k int) int {
    if n == 1 {
        return 0
    }
    return (winnerHelper(n - 1, k) + k) % n
}
