func distanceBetweenBusStops(distance []int, start int, destination int) int {
    i := start
    j := destination
    var cw, ccw int
    
    for start != destination {
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
    
    if cw < ccw {
        return cw
    } else {
        return ccw
    }
}
