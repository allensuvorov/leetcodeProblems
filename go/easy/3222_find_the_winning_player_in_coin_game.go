func losingPlayer(x int, y int) string {
    if min(x, y / 4) % 2 == 1 {
        return "Alice"
    }
    return "Bob"
}
