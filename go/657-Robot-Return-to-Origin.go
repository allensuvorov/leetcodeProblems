func judgeCircle(moves string) bool {
    moveCount := make(map[rune]int)
    for _, m := range moves {
        moveCount[m]++
    }
    return moveCount['U'] == moveCount['D'] && moveCount['R'] == moveCount['L']
}
