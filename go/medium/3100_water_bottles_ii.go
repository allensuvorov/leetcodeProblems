func maxBottlesDrunk(numBottles int, numExchange int) int {
    bottlesDrunk := 0
    emptyBottles := 0

    for numBottles > 0 {
        bottlesDrunk += numBottles // drink all water
        emptyBottles += numBottles // collect all empty bottles
        numBottles = 0
        // exchange
        for emptyBottles >= numExchange {
            emptyBottles -= numExchange // exchange and get 1 bottle
            numBottles++
            numExchange++ // increment
        }
    }

    return bottlesDrunk    
}
