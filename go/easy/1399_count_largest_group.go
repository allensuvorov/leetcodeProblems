func countLargestGroup(n int) int {
    digSumGroupCount := map[int]int{}
    max := 0
    for num := 1; num <= n; num++ {
        temp := num
        sum := 0
        for temp > 0 {
            sum += temp % 10
            temp /= 10
        }
        digSumGroupCount[sum]++
        if digSumGroupCount[sum] > max {
            max = digSumGroupCount[sum]
        }
    }
    countMax := 0
    for _, v := range digSumGroupCount {
        if v == max {
            countMax++
        }
    }
    return countMax
}
