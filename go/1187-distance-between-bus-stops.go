func distanceBetweenBusStops(distance []int, start int, destination int) int {
    // 0 <- start -> n-1
    
    i := start
    cw := 0
    for i != destination {
        cw += distance[i]
        if i == len(distance) - 1 {
            i = 0
        } else {
            i++
        }
    }
    i = start
    ccw := 0
    for i != destination {
        if i == 0 {
            i = len(distance) - 1
        } else {
            i--
        }
        ccw += distance[i]
    }
    if cw < ccw {
        return cw
    } else {
        return ccw
    }
}
