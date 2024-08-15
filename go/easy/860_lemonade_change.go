func lemonadeChange(bills []int) bool {
    cash := make(map[int]int)
    for _, bill := range bills {
        cash[bill]++

        if bill == 5 {
            continue
        }

        if bill == 10 && cash[5] > 0 {
            cash[5]--
            continue
        }

        if bill == 20 && cash[10] > 0 && cash[5] > 0 {
            cash[10]--
            cash[5]--
            continue
        }

        if bill == 20 && cash[5] > 2 {
            cash[5] -= 3
            continue
        }
        
        return false
    }
    return true
}
