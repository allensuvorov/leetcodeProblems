func countPoints(rings string) int {
    res := 0
    rodRings := make([]map[byte]bool, 10)
    for i := 0; i < len(rings); i += 2 {
        color := rings[i]
        rod := rings[i+1] - '0'
        if rodRings[rod] == nil {
            rodRings[rod] = make(map[byte]bool)
        }
        rodRings[rod][color] = true
    }

    for _, ringSet := range rodRings {
        if len(ringSet) == 3 {
            res++
        }
    }
    return res
}
