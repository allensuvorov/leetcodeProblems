func carFleet(target int, position []int, speed []int) int {
    cars := make([][]int, len(position))

    for i := range cars {
        cars[i] = []int{position[i], speed[i]}
    }

    slices.SortFunc(cars, func(a, b []int) int {
        return a[0] - b[0]
    })

    // time of arrival to the target, decreasing stack
    times := []float64{}

    for _, car := range cars {
        time := float64(target - car[0])/float64(car[1])
        for len(times) > 0 && times[len(times)-1] <= time { // catch-up
            times = times[:len(times)-1] // pop
        }
        times = append(times, time)
    }
    return len(times)
}
