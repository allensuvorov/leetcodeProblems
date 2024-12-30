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
func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)
	for _, v := range asteroids {
		stack = append(stack, v)
		if len(stack) > 1 {
			for len(stack) > 1 && stack[len(stack)-2] > 0 && stack[len(stack)-1] < 0 {
				a := stack[len(stack)-2]
				b := stack[len(stack)-1]
				if a < -b {
					stack[len(stack)-2] = b
					stack = stack[:len(stack)-1]
				} else if a > -b {
                    stack = stack[:len(stack)-1]
				} else {
					stack = stack[:len(stack)-2]
                }
			}
		}
	}
    return stack
}
