func distributeCandies(candies int, num_people int) []int {
    ans := make([]int, num_people)
    n := 1
    for candies > 0 {
        for i := range ans {
            if candies < n {
                ans[i] += candies
                return ans
            }
            ans[i] += n
            candies -= n
            n++
        }
    }
    return ans
}
