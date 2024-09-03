func chalkReplacer(chalk []int, k int) int {
    classChalkConsumption := 0
    for i, v := range chalk {
        classChalkConsumption += v
        if classChalkConsumption > k {
            return i
        }
    }

    k = k % classChalkConsumption

    for i, v := range chalk {
        k -= v
        if k < 0 {
            return i
        }
    }
    return 0
}
