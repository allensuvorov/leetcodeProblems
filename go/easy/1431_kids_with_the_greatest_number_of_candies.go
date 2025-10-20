func kidsWithCandies(candies []int, extraCandies int) []bool {
    maxCandies := 0
    for _, v := range candies {
        maxCandies = max(maxCandies, v)
    }

    res := make([]bool, len(candies))
    for i, v := range candies {
        res[i] = v + extraCandies >= maxCandies
    }
    return res
}
