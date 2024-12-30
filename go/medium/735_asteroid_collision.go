func asteroidCollision(asteroids []int) []int {
    stack := make([]int, 0, len(asteroids))
    for _, r := range asteroids {
        top := len(stack) - 1
        // collision cases checks: both exist and both trying to collide
        for len(stack) > 0 && r < 0 && stack[top] > 0 {
            switch {
            case stack[top] > -r: // right asteroid explodes
                r = 0
            case stack[top] < -r: // left asteroid explodes
                stack = stack[:top] 
            default:                       // both explode
                r = 0
                stack = stack[:top]
            }
            top = len(stack) - 1
        }
        if r != 0 { // append, if right asteroid survives collisions
            stack = append(stack, r)
        }
        
    }
    return stack
}

// redo
func removeStars(s string) string {
	stack := []rune{}
	for _, v := range s {
		stack = append(stack, v)
		if len(stack) > 1 {
			a := stack[len(stack)-2]
			b := stack[len(stack)-1]

			for len(stack) > 1 && a > 0 && b < 0 {
				if abs(a) < abs(b) {
					stack[len(stack)-2] = b
					stack = stack[:len(stack)-1]
				} else if abs(a) > abs(b) {
                    stack = stack[:len(stack)-1]
				} else {
					stack = stack[:len(stack)-2]
                }
				a = stack[len(stack)-2]
				b = stack[len(stack)-1]
			}
		}
	}
    return string(stack)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
