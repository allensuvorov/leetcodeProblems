// first idea - start from the end result (reverse engineering)
func isPossible(target []int) bool {
    max := 0
    sum := 0
    maxI:= 0
    
    for {
        for i, num := range target {
            sum += num
            if num > max {
                max = num
                maxI = i
            }
        }

        if max == 1 {
            return true
        }
        if max < sum - max + 1 {
            return false
        } else {
            target[maxI] = sum - max // replace max with dif
        }   
    }
    
}
