func addToArrayForm(num []int, k int) []int {
    ans := []int{}
    digit := 0
    carry := 0
    r := len(num)-1
    
    for k != 0 || r >= 0 {
        digit += carry
        carry = 0

        if r >= 0 {
            digit += num[r]
            r--
        }

        if k > 0 {
            digit += k%10
            k /= 10
        }

        if digit > 9 {
            carry = 1
            digit %= 10
        }

        ans = append([]int{digit}, ans...)
        digit = 0
    }

    if carry == 1 {
        ans = append([]int{carry}, ans...)
    }
    return ans
}
