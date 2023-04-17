package easy

func kidsWithCandies(candies []int, extraCandies int) []bool {
    maxCandies := 0

    for i := range candies {
        if candies[i] > maxCandies {
            maxCandies = candies[i]
        }
    }

    gnc := make([]bool, len(candies))

    for i := range candies {
        if candies[i] + extraCandies >= maxCandies {
            gnc[i] = true
        }
    }
    return gnc
}
