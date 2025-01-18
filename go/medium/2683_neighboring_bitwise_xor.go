func doesValidArrayExist(derived []int) bool {
    first := 0
    curr, next := 0, 0
    for _, v := range derived {
        if v == 1 {
            if curr == 1 {
                next = 0
            } else {
                next = 1
            }
        } else {
            if curr == 1 {
                next = 1
            } else {
                next = 0
            }
        }
        curr = next
    }
    return next == first 
}
