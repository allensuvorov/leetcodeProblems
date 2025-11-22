func matchPlayersAndTrainers(players []int, trainers []int) int {
    // preprocessing
    slices.Sort(players) // n log n
    slices.Sort(trainers) // m log m
    
    j := 0
    result := 0
    for i := range players {
        for j < len(trainers) {
            if players[i] <= trainers[j] {
                result++
                j++
                break
            }
            j++
        }
    }
    return result
}

// logic - for each player we want to find the equal or higher but closest capacity value trainer.

// nested loops
  // outer loop goes over players, one by one, starting from the smallest value
  // and inner loop does a search over the trainers

// this does a linear search, which automatically tracks matched trainers, and reduces the search space.
