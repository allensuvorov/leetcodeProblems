func minimumBoxes(apple []int, capacity []int) int {
    // sort
    slices.Sort(capacity)

    // count apples
    appleCount := 0
    for _, v := range apple {
        appleCount += v
    }

    // start filling boxes starting with bigger ones
    boxCount := 0
    for i := len(capacity) - 1; i >= 0 && appleCount > 0; i-- {
        appleCount -= capacity[i]
        boxCount++
    }

    return boxCount
}


// cap : [1,2,3,4,5]
// apple: [1, 2, 3]
