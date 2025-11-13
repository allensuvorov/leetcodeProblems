func asteroidCollision(asteroids []int) []int {
    stack := []int{}
    for _, v := range asteroids {
        for len(stack) > 0 && v < 0 && stack[len(stack) - 1] > 0 { // collision
            if stack[len(stack) - 1] < -v {
                stack = stack[:len(stack) - 1] // explode
            } else if stack[len(stack) - 1] == -v {
                stack = stack[:len(stack) - 1] // explode
                v = 0 // explode
            } else if stack[len(stack) - 1] > -v {
                v = 0 // explode
            }
        }
        if v != 0 {
            stack = append(stack, v)
        }
    }
    return stack
}
