func kidsWithCandies(candies []int, extraCandies int) []bool {
    gnoc := 0
    for _, v := range candies {
        gnoc = max(gnoc, v)
    }

    res := make([]bool, len(candies))
    for i, v := range candies {
        res[i] = v + extraCandies >= gnoc
    }
    return res
}
