func judgeCircle(moves string) bool {
    moveCount := make(map[rune]int)
    for _, m := range moves {
        moveCount[m]++
    }
    return moveCount['U'] == moveCount['D'] && moveCount['R'] == moveCount['L']
}


// no map/ switch case. much faster.
func judgeCircle(moves string) bool {
    y, x := 0, 0
    for _, m := range moves {
        switch m {
        case 'U': y++
        case 'D': y--
        case 'R': x++
        case 'L': x--
        }
    }
    return x == 0 && y == 0
}
