func squareIsWhite(coordinates string) bool {
    x := coordinates[0] - 'a' + 1
    y := coordinates[1] - '1' + 1
    return (x%2 == 0) != (y % 2 == 0)
}
