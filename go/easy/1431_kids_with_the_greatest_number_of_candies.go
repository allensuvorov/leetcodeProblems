func kidsWithCandies(candies []int, extraCandies int) []bool {
    greatest := 0
    for _, v := range candies {
        greatest = max(greatest, v)
    }

    res := make([]bool,len(candies))
    for i, v := range candies {
        res[i] = v + extraCandies >= greatest
    }
    return res
}
