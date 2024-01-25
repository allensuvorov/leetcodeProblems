func sortColors(nums []int)  {
    cnt := [3]int{}
    for _, v := range nums {
        cnt[v]++
    }
    color := 0
    for i := range nums {
        for cnt[color] == 0 && color < 3 {
            color++
        }
        nums[i] = color
        cnt[color]--
    }
}
