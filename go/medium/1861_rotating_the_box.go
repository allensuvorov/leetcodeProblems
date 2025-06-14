func rotateTheBox(boxGrid [][]byte) [][]byte {
    rows := len(boxGrid[0]) // 3
    cols := len(boxGrid) // 1
    rotatedBox := make([][]byte, rows)
    for r := range rows {
        row := make([]byte, cols)
        for c := range cols {
            row[c] = boxGrid[cols-1-c][r]
        }
        rotatedBox[r] = row
    }

    for c := range cols { // for each column
        for p := rows - 1; p >= 0; p-- { // p - plant index
            if rotatedBox[p][c] != '.' { 
                continue
            }

            // for each empty space
            for s := p - 1; s >= 0; s-- { // s - scout index
                if rotatedBox[s][c] == '*' {
                    p = s
                    break
                }
                if rotatedBox[s][c] == '#' { 
                    rotatedBox[s], rotatedBox[p] = rotatedBox[p], rotatedBox[s]
                    p--
                }
            }
        }
    }
    return rotatedBox
}
