func carFleet(target int, position []int, speed []int) int {
    cars := make([][]int, len(position))

    for i := range cars {
        cars[i] = []int{position[i], speed[i]}
    }

    slices.SortFunc(cars, func(a, b []int) int {
        return a[0] - b[0]
    })

    // time of arrival to the target, decreasing stack
    stack := []float64{}

    for _, cur := range cars {
        curTime := float64(target - cur[0])/float64(cur[1])
        if len(stack) > 0 {
            for len(stack) > 0 && stack[len(stack)-1] <= curTime {
                stack = stack[:len(stack)-1] // pop
            }
        }
        stack = append(stack, curTime)
    }
    return len(stack)
}
