func rotateTheBox(boxGrid [][]byte) [][]byte {
    rows, cols := len(boxGrid[0]), len(boxGrid)
    // rotate the box
    res := make([][]byte, rows)
    for r := range rows {
        res[r] = make([]byte, cols)
        for c := range cols {
            res[r][c] = boxGrid[cols-1-c][r]
        }
    }
   
    for c := range cols { // for each column
        p := rows - 1 // p - plant index begins from bottom    
        for s := rows - 1; s >= 0; s-- { // s - scout index
            if res[s][c] == '#' { 
                res[s][c], res[p][c] = '.', '#'
                p--
            }
            if res[s][c] == '*' {
                p = s - 1
            }
        }
    }
    
    return res
}
