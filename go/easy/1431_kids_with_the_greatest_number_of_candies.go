func kidsWithCandies(candies []int, extraCandies int) []bool {
    gnc := 0
    for _, count := range candies {
        gnc = max(gnc, count)
    }

    res := make([]bool, len(candies))
    for i, count := range candies {
        res[i] = count + extraCandies >= gnc
    }
    return res
}
