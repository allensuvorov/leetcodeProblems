// my solution with nested maps
func isPathCrossing(path string) bool {
    var x, y rune
    track := map[rune]map[rune]bool{0:{0:true}}

    for _, v := range path {
        switch v {
        case 'N': y++
        case 'S': y--
        case 'E': x++
        case 'W': x--
        }
        _, exists := track[x]
        if exists && track[x][y] {
            return true
        } else {
            if exists {
                track[x][y]=true
            } else {
                track[x] = map[rune]bool{y:true}
            }

        }
    }
    return false
}

// top speed solution - uses one map and struct as map key
type Pos struct {
	X, Y int
}

func isPathCrossing(path string) bool {
	var cur Pos
	var cache = map[Pos]struct{}{cur: {}}
	for i := range path {
		switch path[i] {
		case 'N':
			cur.X++
		case 'S':
			cur.X--
		case 'E':
			cur.Y++
		case 'W':
			cur.Y--
		}
		if _, ok := cache[cur]; ok {
			return true
		}
		cache[cur] = struct{}{}
	}
	return false
}
