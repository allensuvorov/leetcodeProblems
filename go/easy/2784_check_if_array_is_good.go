func isGood(nums []int) bool {
    n := len(nums) - 1
    inventory := make(map[int]int)
    
    for i := 1; i < n; i++ {
        inventory[i]++
    }
    inventory[n] = 2
    
    for _, v := range nums {
        inventory[v]--
    }

    for _, v := range inventory {
        if v != 0 {
            return false
        }
    }
    return true
}
