func twoSum(numbers []int, target int) []int {
    l := 0
    r := len(numbers)-1
    for l<r {
        sum := numbers[l] + numbers[r]
        switch{
            case sum == target: return []int{l+1, r+1}
            case sum > target: r--
            case sum < target: l++
        }
    }
    return []int{}
}
