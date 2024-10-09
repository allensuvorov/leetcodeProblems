func averageValue(nums []int) int {
    sum := 0
    cnt := 0
    for _, v := range nums {
        if v % 2 == 0 && v % 3 == 0 {
            sum += v
            cnt++
        }
        fmt.Println(sum, cnt)
    }
    if cnt == 0 {
        return 0
    }
    return sum / cnt
}
