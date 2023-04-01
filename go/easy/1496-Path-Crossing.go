type Position struct {
    x, y int
}

func isPathCrossing(path string) bool {
    var cp Position
    track := map[Position]struct{}{cp:{}}

    for _, v := range path {
        switch v {
        case 'N': cp.y++
        case 'S': cp.y--
        case 'E': cp.x++
        case 'W': cp.x--
        }
        if _, ok := track[cp]; ok {
            return true
        }
        track[cp] = struct{}{}
    }
    return false
}
