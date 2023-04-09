func findMaxK(nums []int) int {

    hm := make(map[int]int)

    for _, v := range nums {

        if _, ok := hm[abs(v)]; ok {
            if hm[abs(v)] + v == 0{
                hm[abs(v)] = 0
            }
        } else {
            hm[abs(v)] = v
        }

    }
    max := -1
    
    for num, v := range hm {
        if v == 0 && num > max {
            max = num
        }
    }
    
    return max
}

func abs (n int) int{
    if n < 0 {
        return n * -1
    }
    return n
}
