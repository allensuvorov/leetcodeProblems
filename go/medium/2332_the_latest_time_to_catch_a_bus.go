func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
    slices.Sort(buses)
    slices.Sort(passengers)
    
    latestTime := 1
    for i, j := 0, 0; i < len(buses); i++ { // simulate the flow of buses
        k := capacity
        seenLastMomentPassenger := false
        // next bus time is buses[i], let's load it up-to K. Any passengers?
        for k > 0 && j < len(passengers) && passengers[j] <= buses[i] {
            // check for the gap
            if j == 0 || passengers[j] - passengers[j - 1] > 1 {
                latestTime = passengers[j] - 1
            }

            if buses[i] == passengers[j] {
                seenLastMomentPassenger = true
            }
            
            j++ // next passenger
            k-- // load the bus
        }

        // bus is not full, could we get on it last moment?
        if k > 0 && !seenLastMomentPassenger { 
            latestTime = buses[i]
        }

    }

    return latestTime
}
