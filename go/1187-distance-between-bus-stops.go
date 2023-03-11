func distanceBetweenBusStops(distance []int, start int, destination int) int {
    i, j := start, destination
    var cw, ccw int
    if start == destination {
        return 0
    }
    for {
        if i == destination {
            if cw < ccw {
                return cw
            }
        } else {
            cw += distance[i] 
            i = (i+1) % len(distance)
        }

        if j == start {
            if ccw < cw {
                return ccw
            }
        } else {
            ccw += distance[j] 
            j = (j+1) % len(distance)
        }
    }
}
